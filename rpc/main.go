package rpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Create an instance of your service implementation
	myService := &MyServiceImpl{}

	// Register your service implementation with the gRPC server
	pb.RegisterMyServiceServer(server, myService)

	// Start the gRPC server and listen on a specific port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server.Serve(listener)
}
