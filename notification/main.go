package main

import (
	"log"
	"net"
	"remind/notification/server"
	"remind/pkg/logger"

	pb "remind/notification/pb"

	"remind/notification/database"
	"remind/notification/firebase"

	"google.golang.org/grpc"
)

func main() {

	firebase.Init()
	print("Firebase initialized")
	database.Init()
	// Listen on a TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server started on port localhost:50051")
	// Create a new gRPC server
	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)

	s := grpc.NewServer(grpcLogger)
	// Register your server with the gRPC server
	pb.RegisterNotificationServiceServer(s, &server.Server{})

	log.Println("Server registered")
	// Start the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
