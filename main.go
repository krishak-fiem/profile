package main

import (
	"github.com/krishak-fiem/db/go/cassandra"
	"github.com/krishak-fiem/profile/kafka"
	"github.com/krishak-fiem/profile/proto/pb"
	"github.com/krishak-fiem/profile/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")
	cassandra.Init(9043)
	go kafka.UserCreatedReader()
	if err != nil {
		log.Fatalf("Failed to start the auth products on port 5000: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	profile := service.Server{}

	pb.RegisterProfileServiceServer(grpcServer, &profile)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the auth products on port 5000: %v\n", err)
	}
}
 