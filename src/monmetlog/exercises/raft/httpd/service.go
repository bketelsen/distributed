// Package httpd provides the HTTP server for accessing the distributed key-value store.
// It also provides the endpoint for other nodes to join an existing cluster.
package httpd

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

// Service provides HTTP service.
type Service struct {
	addr string
	ln   net.Listener

	// Store is the interface Raft-backed key-value stores must implement.
	Store interface {
		// Get returns the value for the given key, and ensures a consensus read.
		Get(key string) (string, error)

		// Set sets the value for the given key, via distributed consensus.
		Set(key, value string) error

		// Delete removes the given key, via distributed consensus.
		Delete(key string) error

		// AddPeer adds the node to the cluster.
		AddPeer(addr string) error

		// Leader will return the current leader of the cluster
		Leader() string
	}
}

// New returns an uninitialized HTTP service.
func New(addr string) *Service {
	return &Service{
		addr: addr,
	}
}

// Start starts the service.
func (s *Service) Start() error {
	server := http.Server{
		Handler: s,
	}

	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.ln = ln

	http.Handle("/", s)

	go func() {
		err := server.Serve(s.ln)
		if err != nil {
			log.Fatalf("HTTP serve: %s", err)
		}
	}()

	return nil
}

// Close closes the service.
func (s *Service) Close() {
	s.ln.Close()
	return
}

// ServeHTTP allows Service to serve HTTP requests.
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/key") {
		s.handleKeyRequest(w, r)
	} else if r.URL.Path == "/join" {
		s.handleJoin(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *Service) handleJoin(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(m) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	remoteAddr, ok := m["addr"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.Store.AddPeer(remoteAddr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func parseKey(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return ""
	}
	return parts[2]
}

func (s *Service) handleKeyRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		k := parseKey(r.URL.Path)
		if k == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		v, err := s.Store.Get(k)
		if err != nil {
			w.Header().Set("X-RAFT-LEADER", s.Store.Leader())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(map[string]string{k: v})
		if err != nil {
			w.Header().Set("X-RAFT-LEADER", s.Store.Leader())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		io.WriteString(w, string(b))

	case "POST":
		// Read the value from the POST body.
		m := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for k, v := range m {
			if err := s.Store.Set(k, v); err != nil {
				w.Header().Set("X-RAFT-LEADER", s.Store.Leader())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	case "DELETE":
		k := parseKey(r.URL.Path)
		if k == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		if err := s.Store.Delete(k); err != nil {
			w.Header().Set("X-RAFT-LEADER", s.Store.Leader())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.Store.Delete(k)

	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
	return
}
