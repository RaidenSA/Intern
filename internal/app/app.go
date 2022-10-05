package app

import (
	"google.golang.org/grpc"
	"intern/internal/api"
	"intern/internal/storage"
	"log"
	"net"
	"net/http"
)

type Server struct {
	Storage storage.MemoryStorage
}

func New(storageName string, serverMode string) {
	var stor storage.MemoryStorage
	switch storageName {
	case "postgres":
		log.Println("Storage mode: ", storageName)
		stor = storage.DataBase{
			ConnStr: "user=postgres password=fnkfynblf dbname=dbase sslmode=disable",
		}
	default:
		log.Println("Storage mode: InMemory")
		stor = &storage.Container{
			MapTokenToValue: make(map[string]string),
			MapValueToToken: make(map[string]string),
		}
	}

	s := &Server{
		Storage: stor,
	}
	switch serverMode {
	case "grpc":
		log.Println("Server mode: ", serverMode)
		serv := grpc.NewServer()
		// Register gRPC server
		api.RegisterPostListenerServer(serv, s)

		// Listen on port 8080
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal(err)
		}

		// Start gRPC server
		if err := serv.Serve(l); err != nil {
			log.Fatal(err)
		}
	default:
		log.Println("Server mode: HTTP")
		http.HandleFunc("/", s.Handler)
		log.Fatal(http.ListenAndServe(Addr, nil))
	}
	return
}
