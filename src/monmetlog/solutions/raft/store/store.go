package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

const (
	// how long should we wait to get a consensus read
	barrierTimeout = 5 * time.Second
	// how many snapshots to retain
	retainSnapshotCount = 2
	// how long should we wait for raft in general to time out an operation
	raftTimeout = 10 * time.Second
)

var ErrNotLeader = errors.New("not leader")

type Store struct {
	path     string // where to store the raft database
	bindAddr string // what address to bind the raft service on

	raft *raft.Raft // The raft instance
	fsm  *fsm       // our finite state machine

	logger *log.Logger // Use your own logger.
}

// New returns an instance of a new store
func New() *Store {
	return &Store{
		logger: log.New(os.Stderr, "[raft-store] ", log.LstdFlags),
		fsm: &fsm{
			data: make(map[string]string),
		},
	}
}

// Open will attempt to start up raft and connect to the consensus cluster
// If we detect we have no peers, and haven't been given a join address,
// we will attempt to bootstrap and become the leader
func (s *Store) Open(joinAddr, path, bindAddr string) error {
	// store arguments
	s.path = path
	s.bindAddr = bindAddr

	// Setup Raft configuration.
	config := raft.DefaultConfig()

	// Retrieve the peers on disk.  Raft will write this file automatically.
	peers, err := retrievePeers(filepath.Join(s.path, "peers.json"))
	if err != nil {
		return err
	}

	// Allow the node to entry single-mode, potentially electing itself, if
	// no join address specified and we are the only node.
	if joinAddr == "" && len(peers) <= 1 {
		s.logger.Println("enabling single-node mode")
		config.EnableSingleNode = true
		config.DisableBootstrapAfterElect = false
	}

	// Setup Raft communication.
	addr, err := net.ResolveTCPAddr("tcp", s.bindAddr)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(s.bindAddr, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	// Create peer storage.
	peerStore := raft.NewJSONPeers(s.path, transport)

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(s.path, retainSnapshotCount, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}

	// Create the log store and stable store.
	logStore, err := raftboltdb.NewBoltStore(filepath.Join(s.path, "raft.db"))
	if err != nil {
		return fmt.Errorf("new bolt store: %s", err)
	}

	// Instantiate the Raft systems.
	ra, err := raft.NewRaft(config, s.fsm, logStore, logStore, snapshots, peerStore, transport)
	if err != nil {
		return fmt.Errorf("new raft: %s", err)
	}
	s.raft = ra

	s.logger.Println("started successfully")

	return nil
}

// Get returns the value for the given key.
// You must ONLY read from the leader to ensure you have a consensus read
// You must also make sure all applied writes have been written before reading
func (s *Store) Get(key string) (string, error) {
	// Not the leader, you can't ensure a consenus read
	if s.raft.State() != raft.Leader {
		return "", ErrNotLeader
	}

	// Make sure we apply all writes before returning
	f := s.raft.Barrier(barrierTimeout)
	if f.Error() != nil {
		return "", f.Error()
	}

	// Finally, it's safe to assume we have the most up to date information
	return s.fsm.Get(key), nil
}

// Set sets the value for the given key.
func (s *Store) Set(key, value string) error {
	if s.raft.State() != raft.Leader {
		return ErrNotLeader
	}

	c := &command{
		Command: "set",
		Key:     key,
		Value:   value,
	}
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	f := s.raft.Apply(b, raftTimeout)
	return f.Error()
}

// Delete deletes the given key.
func (s *Store) Delete(key string) error {
	if s.raft.State() != raft.Leader {
		return ErrNotLeader
	}

	c := &command{
		Command: "delete",
		Key:     key,
	}
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	f := s.raft.Apply(b, raftTimeout)
	return f.Error()
}

// AddPeer will attempt to join this node to an existing cluster
func (s *Store) AddPeer(addr string) error {
	s.logger.Printf("attempting to join %s", addr)

	f := s.raft.AddPeer(addr)
	if f.Error() != nil {
		return f.Error()
	}
	s.logger.Printf("joined %s", addr)
	return nil
}

// Leader will return the current leader of the cluster
func (s *Store) Leader() string {
	return s.raft.Leader()
}

func retrievePeers(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)

	// If the file doesn't exist, we have no peers.
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// If the file is empty, we have no peers
	if len(b) == 0 {
		return nil, nil
	}

	var peers []string
	dec := json.NewDecoder(bytes.NewReader(b))
	if err := dec.Decode(&peers); err != nil {
		return nil, err
	}

	return peers, nil
}
