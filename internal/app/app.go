package app

import (
	"intern/internal/handler"
	"intern/internal/service"
	"intern/internal/storage"
	"net/http"
)

var Server = &http.Server{
	Addr: handler.Addr,
}

func New() {
	service.ServiceStorage.CurStorage = &storage.Container{
		MapTokenToValue: make(map[string]string),
		MapValueToToken: make(map[string]string),
	}
	//service.ServiceStorage.CurStorage = storage.DataBase{}
	http.HandleFunc("/", handler.HTTPHandler)
	return
}
