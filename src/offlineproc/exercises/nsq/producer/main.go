package main

import (
	"flag"
	"log"

	nsq "github.com/bitly/go-nsq"
)

func main() {

	var website string
	nsqAddr := "127.0.0.1:4150"
	flag.StringVar(&website, "website", "", "provide a website to add to the ping que")
	flag.StringVar(&nsqAddr, "nsqaddr", nsqAddr, "provide nsq addr to connect to")
	flag.Parse()
	if website == "" {
		log.Println("no website provided")
		flag.Usage()
		return
	}
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(nsqAddr, config)

	err := w.Publish("ping", []byte(website))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()

}
