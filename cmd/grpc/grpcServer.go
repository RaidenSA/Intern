package main

import (
	"google.golang.org/grpc"
	"intern/internal/api"
	"intern/internal/app"
	"log"
	"net"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &app.Server{}

	// Register gRPC server
	api.RegisterPostListenerServer(s, srv)

	// Listen on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
