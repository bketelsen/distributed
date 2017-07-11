package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"offlineproc/exercises/nsq/client/consumer"

	nsq "github.com/bitly/go-nsq"
)

func main() {
	path := "/tmp/nsqoutput"

	flag.StringVar(&path, "path", path, "path to store output of nsq consumers")
	flag.Parse()

	if err := os.MkdirAll(path, 0700); err != nil {
		log.Println(err)
		return
	}

	config := nsq.NewConfig()
	service := consumer.New(consumer.DefaultNSQDLookupHost, path, config)

	if err := service.Open(); err != nil {
		log.Println("failed to open consumer service")
		log.Fatal(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	log.Println("interrupt signal received, terminating")
	service.Close()
}
