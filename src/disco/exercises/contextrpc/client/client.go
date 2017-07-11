package main

import (
	"bytes"
	"disco/exercises/contextrpc"
	"log"
	"net/http"

	"github.com/gorilla/rpc/json"
)

func main() {

	url := "http://localhost:1234/rpc"
	args := &contextrpc.HelloArgs{
		Who: "Gopher",
	}
	message, err := json.EncodeClientRequest("HelloService.Say", args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	var result contextrpc.HelloReply
	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("%s\n", result.Message)
}
