// Copyright 2012-2016 Apcera Inc. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/nats-io/nats"
)

func usage() {
	log.Fatalf("Usage: publisher [-s server (%s)] <subject> \n", nats.DefaultURL)
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	fmt.Printf("Connecting to %s\n", *urls)
	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	for x := 1; ; x++ {
		subj, msg := args[0], []byte(strconv.Itoa(x))

		nc.Publish(subj, msg)
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'\n", subj, msg)
		}
	}
}
