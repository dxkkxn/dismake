package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"bufio"

	"google.golang.org/grpc"
	pb "dismake/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCommandRemoteExecServer
}

func (s *server) CmdRemoteExec(ctx context.Context, in *pb.CmdRequest) (*pb.CmdResponse, error) {
	log.Printf("received: %v\n", in.GetCmd())
	log.Printf("excuting command %v", in.GetCmd())

	cmd := exec.Command("bash", "-c", in.GetCmd())
	stderr, _ := cmd.StderrPipe()
	var res string;
	if err := cmd.Start(); err != nil {
		log.Fatal("ERROR: %v", err)
	}
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		error := scanner.Text()
		if error != "" {
			log.Fatal("CMD ERROR: %v", error)
		}
	}
	return &pb.CmdResponse{Res: res}, nil
}

func main() {
	flag.Parse()
	hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
	// fmt.Printf("Hostname: %s", hostname)
	log.SetPrefix("[server: " + hostname + "]")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommandRemoteExecServer(s, &server{})
	log.Printf(" listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
