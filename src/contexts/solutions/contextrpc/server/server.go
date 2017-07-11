package main

import (
	"contexts/solutions/contextrpc"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *contextrpc.HelloArgs, reply *contextrpc.HelloReply) error {
	time.Sleep(600 * time.Millisecond)
	reply.Message = "Hello, " + args.Who + "!"

	return nil
}
func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	http.Handle("/rpc", s)
	fmt.Println("Listening")
	http.ListenAndServe(":1234", nil)
}
