package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	domain "github.com/ssouthcity/go-grpc-first-look"
	"github.com/ssouthcity/go-grpc-first-look/proto"
)

func main() {
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	var carSrv domain.CarService
	carSrv = proto.NewCarClient(conn)

	car, err := carSrv.GetCar()
	if err != nil {
		log.Fatalf("failed to get car: %v", err)
	}

	fmt.Println(car)
}
