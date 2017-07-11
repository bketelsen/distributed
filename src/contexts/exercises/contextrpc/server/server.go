package main

import (
	"contexts/exercises/contextrpc"
	"fmt"
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *contextrpc.HelloArgs, reply *contextrpc.HelloReply) error {
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
