package main

import (
	"Intern/internal/handler"
	"log"
	"net/http"
)

const addr = "localhost:8080"

func main() {
	// In-Memory storage
	//Selection of handler
	http.HandleFunc("/", handler.TemplateHandler)
	//Server Properties
	server := &http.Server{
		Addr: addr,
	}
	log.Fatal(server.ListenAndServe())

}
