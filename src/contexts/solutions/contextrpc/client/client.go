package main

import (
	"bytes"
	"context"
	"contexts/solutions/contextrpc"
	"log"
	"net/http"
	"time"

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
	var cf context.CancelFunc
	ctx, cf := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cf()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	f := func(resp *http.Response, e error) error {
		if e != nil {
			return e
		}
		defer resp.Body.Close()

		var result contextrpc.HelloReply
		err = json.DecodeClientResponse(resp.Body, &result)
		if err != nil {
			log.Fatalf("Couldn't decode response. %s", err)
			return err
		}
		log.Printf("%s\n", result.Message)
		return nil
	}
	err = httpDo(ctx, req, f)
	if err != nil {
		log.Println(err)
	}

}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() {
		c <- f(client.Do(req))
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
