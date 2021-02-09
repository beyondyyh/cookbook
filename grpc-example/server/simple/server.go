package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "grpc-example/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Println(req.String(), req.GetPageSize())

	var bundles = make(map[string]*pb.Bundle)
	for i := 0; i < 3; i++ {
		bundles[strconv.Itoa(i)] = &pb.Bundle{
			AppId:   "app_id" + strconv.Itoa(i),
			AppName: "app_name" + strconv.Itoa(i),
			Url:     "url" + strconv.Itoa(i),
		}
	}
	return &pb.SearchResponse{
		Bundles: bundles,
	}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen error: %v", err)
	}
	server.Serve(lis)
}
