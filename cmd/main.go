package main

import (
	"google.golang.org/grpc"
	"intern/internal/api"
	"intern/internal/app"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// Here should be params
	args := os.Args[1:]

	storageMode := "inMemory"
	if len(args) > 0 {
		storageMode = args[0]
	}
	s := app.New(storageMode) //"postgres for postgres mode

	go func(srv *app.Server) {
		s := grpc.NewServer()
		api.RegisterPostListenerServer(s, srv)
		// Listen on port 8080
		l, err := net.Listen("tcp", ":8088")
		if err != nil {
			log.Fatal(err)
		}
		// Start gRPC server
		log.Println("Grpc Running")
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}(s)
	log.Println("HTTP Runing")
	log.Fatal(http.ListenAndServe(app.Addr, nil))
}
