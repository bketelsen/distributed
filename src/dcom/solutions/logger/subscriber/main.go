// Copyright 2012-2016 Apcera Inc. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/nats-io/nats"
)

// NOTE: Use tls scheme for TLS, e.g. nats-sub -s tls://demo.nats.io:4443 foo
func usage() {
	fmt.Println("Usage: nats-sub [-s server] [-t] <subject> \n")
}

func printMsg(f *os.File, m *nats.Msg, i int) {
	msg := fmt.Sprintf("[#%d] [%s]: '%s'\n", i, m.Subject, string(m.Data))
	_, err := f.WriteString(msg)
	if err != nil {
		log.Println("Error writing to file: ", err)
	}
	err = f.Sync()
	if err != nil {
		log.Println("Error syncing file: ", err)
	}
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var showTime = flag.Bool("t", false, "Display timestamps")
	f, err := os.Create("nats.log")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	fmt.Printf("Connecting to %s\n", *urls)
	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	subj, i := args[0], 0

	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		printMsg(f, msg, i)
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]\n", subj)
	if *showTime {
		log.SetFlags(log.LstdFlags)
	}

	runtime.Goexit()
	defer f.Close()
}
