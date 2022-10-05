package main

import (
	"intern/internal/app"
	"os"
)

func main() {
	// Here should be params
	args := os.Args[1:]

	storageMode := "inMemory"
	if len(args) > 0 {
		storageMode = args[0]
	}
	servermode := "HTTP"
	if len(args) > 1 {
		servermode = args[1]
	}
	app.New(storageMode, servermode) //"postgres for postgres mode

}
