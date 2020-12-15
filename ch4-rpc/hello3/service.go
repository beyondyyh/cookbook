package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "foo/hello3/HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
	// Foo(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	fmt.Printf("Deal request %s\n", request)
	*reply = "Hello " + request
	return nil
}

func (s *HelloService) Foo(request string, reply *string) error {
	fmt.Printf("Deal request %s\n", request)
	*reply = "Welcome " + request
	return nil
}

func main() {
	// svc := &HelloService{}
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
