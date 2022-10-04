package main

import (
	"intern/internal/app"
	"intern/internal/handler"
	"log"
	"net/http"
)

const addr = "localhost:8080"

func main() {
	// In-Memory storage
	//Selection of handler
	app.New()
	http.HandleFunc("/", handler.TemplateHandler)
	//Server Properties
	server := &http.Server{
		Addr: addr,
	}
	log.Fatal(server.ListenAndServe())

}
