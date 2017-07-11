package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"monmetlog/solutions/raft/httpd"
	"monmetlog/solutions/raft/store"

	"github.com/bketelsen/trace"
)

const (
	defaultHTTPAddr = ":8080"
	defaultRaftAddr = ":8086"
)

// Command line parameters
var httpAddr string
var raftAddr string
var joinAddr string

func init() {
	flag.StringVar(&httpAddr, "httpaddr", defaultHTTPAddr, "Set the HTTP bind address")
	flag.StringVar(&raftAddr, "raftaddr", defaultRaftAddr, "Set Raft bind address")
	flag.StringVar(&joinAddr, "join", "", "[optional] The address of a node to join.  Leave empty to boostrap your first node.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <raft-data-path> \n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Usage = usage
}

func main() {
	eventLog := trace.NewEventLog("raft", "service")
	defer eventLog.Finish()
	// separate handler for /debug/* requests on trace server
	go http.ListenAndServe("127.0.0.1:6060", nil)
	go func() {
		for {
			result, err := trace.DumpMetrics(context.Background(), "")
			if err != nil {
				eventLog.Errorf("Error dumping metrics %#v", err)
			}
			fmt.Println("======================================")
			fmt.Println(result)
			fmt.Println("======================================\n\n")
			time.Sleep(5 * time.Second)
		}
	}()

	l, err := net.Listen("tcp", ":6061")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to listen on metrics port\n")
		eventLog.Errorf("Unable to listen on metrics port")
		os.Exit(1)
	}
	go trace.ServeMetrics(context.Background(), l)
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		eventLog.Errorf("No Raft storage directory specified")
		os.Exit(1)
	}

	// Ensure Raft storage exists.
	raftDir := flag.Arg(0)
	if raftDir == "" {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		eventLog.Errorf("No Raft storage directory specified")
		os.Exit(1)
	}
	os.MkdirAll(raftDir, 0700)

	s := store.New()
	if err := s.Open(joinAddr, raftDir, raftAddr); err != nil {
		eventLog.Errorf("failed to open store: %s", err.Error())
		log.Fatalf("failed to open store: %s", err.Error())
	}

	h := httpd.New(httpAddr)
	// Give the instance of our httpd service a reference to the raft store
	h.Store = s
	eventLog.Printf("Starting HTTP Service")
	if err := h.Start(); err != nil {
		log.Fatalf("failed to start HTTP service: %s", err.Error())
	}
	log.Printf("started http service on %s", httpAddr)

	eventLog.Printf("Started HTTP Service on %s", httpAddr)
	// If join was specified, make the join request.
	if joinAddr != "" {
		if err := join(joinAddr, raftAddr); err != nil {
			eventLog.Errorf("failed to join node at %s: %s", joinAddr, err.Error())
			log.Fatalf("failed to join node at %s: %s", joinAddr, err.Error())
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	eventLog.Printf("interrupt signal received, terminating")
	log.Println("interrupt signal received, terminating")
}

// Since you can only add a peer if you are the leader
// we need to actually post a request to this address
// and request we join
func join(joinAddr, raftAddr string) error {
	b, err := json.Marshal(map[string]string{"addr": raftAddr})
	if err != nil {
		return err
	}
	host := fmt.Sprintf("http://%s/join", joinAddr)

	log.Printf("attempting to join %s", host)
	resp, err := http.Post(host, "application-type/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func usage() {
	fmt.Println(`
usage: raft [options] path

	path: path on disk to store your raft database and peers.json

Start your first node:

raft -httpaddr localhost:8180 --raftaddr localhost:8186 /tmp/raft1

Start your second node:

	raft --httpaddr localhost:8280 --raftaddr localhost:8286 --join localhost:8180 /tmp/raft2

Start your third node:

	raft --httpaddr localhost:8380 --raftaddr localhost:8386 --join localhost:8180 /tmp/raft3


Options:
	`)
	flag.PrintDefaults()
}
