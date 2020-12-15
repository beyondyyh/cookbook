package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 调试看效果：
// 1. 运行服务端 go run service.go
// 2. 新开tab运行：$ echo -e '{"foo/hello4/HelloService.Hello","params":["nc demo"],"id":1}' | nc localhost 1234
// 3. service会Accept请求，输出结果：{"id":1,"result":"Hello nc demo","error":null}

const HelloServiceName = "foo/hello4/HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
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

func main() {
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

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
