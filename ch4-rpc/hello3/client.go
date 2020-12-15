package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

// var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DailHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func (p *HelloServiceClient) Foo(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Foo", request, reply)
}

func helloClient() {
	client, err := DailHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dailing:", err)
	}

	var reply string
	err = client.Hello("hello3", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

	err = client.Foo("foo", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
