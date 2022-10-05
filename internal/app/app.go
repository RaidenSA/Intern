package app

import (
	"intern/internal/storage"
	"net/http"
)

type Server struct {
	Storage storage.MemoryStorage
}

func New(storageName string) *Server {
	var stor storage.MemoryStorage
	switch storageName {
	case "postgres":
		stor = storage.DataBase{
			ConnStr: "user=postgres password=fnkfynblf dbname=dbase sslmode=disable",
		}
	default:
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
