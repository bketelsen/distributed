package main

import (
	"disco/exercises/contextrpc"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/hashicorp/consul/api"
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
	log.Println("Registering with Consul")

	ip := os.Getenv("SERVER_IP")
	if ip == "" {
		log.Fatal("Must set SERVER_IP Environment Variable.")
	}
	service := &api.AgentServiceRegistration{}
	service.Name = "hello"
	service.Address = ip
	service.Port = 9876
	conf := api.Config{}
	conf.Address = "127.0.0.1:8500"
	client, err := api.NewClient(&conf)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceRegister(service)
	if err != nil {
		panic(err)
	}
	log.Println("Listening")
	http.ListenAndServe(":9876", nil)
}
