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

	var query string = "grpc-example"
	var pageno int32 = 1
	var pagesize int32 = 20
	// var corpus = pb.SearchRequest2_IMAGES
	client := pb.NewSearchService2Client(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest2{
		Query:         &query,
		PageNumber:    &pageno,
		ResultPerPage: &pagesize,
		// Corpus:        &corpus,
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	b, _ := json.Marshal(resp.GetBundles())
	log.Printf("resp: %s", string(b))
}
