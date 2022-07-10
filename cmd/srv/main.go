package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ssouthcity/go-grpc-first-look/proto"
	"github.com/ssouthcity/go-grpc-first-look/proto/pb"
)

func main() {
	lst, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCarServiceServer(s, &proto.CarServer{})

	if err := s.Serve(lst); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
