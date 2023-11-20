package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "pingpong/pingpong"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedPingPongServer
}

func (s *server) Pong(ctx context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	// log.Println("[server] received ping")
	// log.Println("[server] sending pong")
	// log.Printf("[server] %v\n", in.GetMessage())

	return &pb.PongResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingPongServer(s, &server{})
	log.Printf("[server] listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
