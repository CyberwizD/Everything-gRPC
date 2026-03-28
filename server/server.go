package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Listen using the TCP on port 9000
	port, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen, %v", err)
	}
	
	// Initialize gRPC server
	grpcServer := grpc.NewServer()

	// Serve gRPC server on port 9000
	if err := grpcServer.Serve(port);
	err != nil {
		log.Fatalf("Failed to serve gRPC comm, %v", err)
	}
}

