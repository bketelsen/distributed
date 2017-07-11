// Copyright 2012-2016 Apcera Inc. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nats-io/nats"
)

func usage() {
	log.Fatalf("Usage: publisher [-s server (%s)] <subject> <msg> \n", nats.DefaultURL)
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usage()
	}
	fmt.Printf("Connecting to %s\n", *urls)
	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subj, msg := args[0], []byte(args[1])

	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
