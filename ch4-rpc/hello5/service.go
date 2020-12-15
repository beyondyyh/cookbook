package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "foo/hello5/HelloService"

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

// Go内置Http协议上的RPC框架
// 如何测试？
// 1. go run service.go
// 2. curl localhost:1234/jsonrpc -X POST --data '{"method":"foo/hello5/HelloService.Hello","params":["http demo"],"id":0}'
// output: {"id":0,"result":"Hello http demo","error":null}
func main() {
	RegisterHelloService(new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
	select {}
}
