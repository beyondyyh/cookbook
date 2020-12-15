package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func helloClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dailing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "helloClient", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
