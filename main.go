package main

import (
	"log"
	"net/http"
	"os"
)

var db Storage
var server *http.Server
var logger *log.Logger

func init() {

	db = NewStorage()

	logger = log.New(os.Stdout, "", 0)

	server = Start(logger)

	logger.Printf("Listening on http://0.0.0.0%s\n", server.Addr)
}

func main() {

	log.Fatal(server.ListenAndServe())
}
