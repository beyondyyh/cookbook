package main

// client_pub 从客户端向服务器发布消息

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/beyondyyh/cookbook/grpc-pubsub/mypubsub"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := mypubsub.NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(), &mypubsub.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &mypubsub.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
