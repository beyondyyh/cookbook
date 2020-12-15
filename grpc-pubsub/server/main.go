package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"

	"github.com/beyondyyh/cookbook/grpc-pubsub/mypubsub"
)

// Wrap mypubsub.String
type String = mypubsub.String

type PubsubService struct {
	pub *pubsub.Publisher
}

// NewPubsubService
func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

// 发布
func (p *PubsubService) Publish(ctx context.Context, arg *String) (*String, error) {
	p.pub.Publish(arg.GetValue())
	return &String{}, nil
}

// 订阅
func (p *PubsubService) Subscribe(arg *String, stream mypubsub.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

// Server, gRPC服务的启动流程
func main() {
	grpcServer := grpc.NewServer()
	mypubsub.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}
