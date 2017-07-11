package main

import (
	"bytes"
	"disco/solutions/contextrpc"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/rpc/json"
	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	services, _, err := client.Catalog().Service("hello", "", nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, service := range services {
		fmt.Printf("Service %s, listening on %s:%d \n", service.ServiceName, service.ServiceAddress, service.ServicePort)
	}
	if len(services) < 1 {
		log.Fatal("No services available")
	}

	url := fmt.Sprintf("http://%s:%d/rpc", services[0].ServiceAddress, services[0].ServicePort)
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
	cl := new(http.Client)
	resp, err := cl.Do(req)
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
