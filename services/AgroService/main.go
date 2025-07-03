package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/longkhan786/dhareja-platform/agroservice/gen"
	"github.com/longkhan786/dhareja-platform/agroservice/server"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAgroTradersServiceServer(grpcServer, &server.AgroService{})

	reflection.Register(grpcServer)

	log.Println("AgroService is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
