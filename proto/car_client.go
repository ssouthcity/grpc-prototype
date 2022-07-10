package proto

import (
	"context"

	domain "github.com/ssouthcity/go-grpc-first-look"
	"github.com/ssouthcity/go-grpc-first-look/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CarClient struct {
	srv pb.CarServiceClient
}

func NewCarClient(conn *grpc.ClientConn) *CarClient {
	return &CarClient{pb.NewCarServiceClient(conn)}
}

func (c *CarClient) GetCar() (*domain.Car, error) {
	car, err := c.srv.GetCar(context.TODO(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return c.mapCar(car), nil
}

func (c *CarClient) mapCar(car *pb.Car) *domain.Car {
	return &domain.Car{
		Brand: car.Brand,
		Model: car.Model,
	}
}
