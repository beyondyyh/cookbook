package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-example/proto"
)

type StreamService struct{}

const PORT = "9002"

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen error: %v", err)
	}
	server.Serve(lis)
}

// 客户端发起一次普通的RPC请求，服务端通过流式响应多次发送数据集，客户端Recv接收
func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	log.Printf("request: %s", r.String())
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 客户端通过流式发送多次RPC请求给服务端，服务端发起一次响应给客户端
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{
				Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1},
			})
		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name:%s, pt.value:%d", r.Pt.Name, r.Pt.Value)
	}
}

// 双向流式RPC，由客户端以流式方式发起请求，服务端同样以流式的方式响应请求，假设双向流是按顺序发送的
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gRPC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++

		log.Printf("stream.Recv pt.name:%s, pt.value:%d", r.Pt.Name, r.Pt.Value)
	}
}
