package main

import (
	"intern/internal/app"
	"log"
	"net/http"
)

func main() {
	// Here should be params

	app.New("postgres") //"postgres for postgres mode
	log.Fatal(http.ListenAndServe(app.Addr, nil))

}
