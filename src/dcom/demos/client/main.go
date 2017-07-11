package main

import (
	"fmt"
	"log"
	"net/rpc"

	"dcom/demos/server/stringsvc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &stringsvc.Args{Name: "Brian"}

	var reply stringsvc.Result
	err = client.Call("Upper.Uppercase", args, &reply)
	if err != nil {
		log.Fatal("stringsvc error:", err)
	}
	fmt.Printf("Synchronous Call to Upper: %s=%s\n", args.Name, reply.Name)

	// Asynchronous call
	result := new(stringsvc.Result)
	upperCall := client.Go("Upper.Uppercase", args, result, nil)
	replyCall := <-upperCall.Done
	if replyCall.Error != nil {
		log.Fatal("async stringsvc error:", err)
	}

	fmt.Printf("Asynchronous Call to Upper: %s=%s\n", args.Name, result.Name)
}
