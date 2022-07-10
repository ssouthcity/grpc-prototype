package proto

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ssouthcity/go-grpc-first-look/proto/pb"
)

type CarServer struct {
	pb.UnimplementedCarServiceServer
}

func (s *CarServer) GetCar(context.Context, *empty.Empty) (*pb.Car, error) {
	return &pb.Car{Brand: "BMW", Model: "i3"}, nil
}
