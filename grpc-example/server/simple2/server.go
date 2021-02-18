package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "grpc-example/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, req *pb.SearchRequest2) (*pb.SearchResponse2, error) {
	log.Println(req.String(), req.GetCorpus())
	if req.GetCorpus() == pb.SearchRequest2_IMAGES {
		fmt.Println("corp matches")
	}

	var bundles = make([]*pb.Bundle, 0)
	for i := 0; i < 3; i++ {
		bundles = append(bundles, &pb.Bundle{
			AppId:   "app_id" + strconv.Itoa(i),
			AppName: "app_name" + strconv.Itoa(i),
			Url:     "url" + strconv.Itoa(i),
		})
	}
	return &pb.SearchResponse2{
		Bundles: bundles,
	}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchService2Server(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen error: %v", err)
	}
	server.Serve(lis)
}
