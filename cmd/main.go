package main

import (
	"intern/internal/app"
	"log"
)

func main() {
	// Here should be params
	app.New()
	log.Fatal(app.Server.ListenAndServe())

}
