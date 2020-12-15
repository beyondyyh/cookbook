package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func getClient() *rpc.Client {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dail:", err)
	}

	return rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call(KVStoreServiceName+".Watch", 5, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()

	if err := client.Call(
		KVStoreServiceName+".Set", [2]string{"abc", "abc-value11"},
		new(struct{}),
	); err != nil {
		log.Fatal(err)
	}

	if err := client.Call(
		KVStoreServiceName+".Set", [2]string{"xyz", "xyz-value11"},
		new(struct{}),
	); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 2)
}
