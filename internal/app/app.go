package app

import (
	"intern/internal/storage"
	"log"
)

const user = "postgres"
const myPass = "postgres"
const dbname = "postgres"
const connStr = "host=postgres" + " port=5432" + " user=" + user + " password=" + myPass + " dbname=" + dbname + " sslmode=disable"

type Server struct {
	Storage storage.MemoryStorage
}

func New(storageName string) *Server {
	var stor storage.MemoryStorage
	//Selecting storage depending on argument
	switch storageName {
	case "postgres":
		log.Println("Storage mode: ", storageName)
		stor = storage.DataBase{
			ConnStr: connStr,
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

	return s
}
