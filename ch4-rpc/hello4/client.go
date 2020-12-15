package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 调试看效果：
// 1. 用nc启动一个tcp服务监听1234端口：nc -l 1234
// 2. 新开tab运行客户端：go test -v -run Test_helloClient
// 3. nc server端会输出：{"method":"foo/hello4/HelloService.Hello","params":["client4"],"id":0}

func helloClient() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dail:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(HelloServiceName+".Hello", "client4", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
