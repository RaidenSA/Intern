package app

import (
	"log"
	"net/http"

	"intern/internal/storage"
)

type Server struct {
	Storage storage.MemoryStorage
}

func New(storageName string) *Server {
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

	http.HandleFunc("/", s.Handler)
	return s
}
