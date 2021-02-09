package main

import (
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc"

	pb "grpc-example/proto"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dail err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Query:  "google",
		PageNo: 1,
		// PageSize: 50,
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	b, _ := json.Marshal(resp.GetBundles())
	log.Printf("resp: %s", string(b))
}
