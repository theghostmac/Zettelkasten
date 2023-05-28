package main

import (
	"log"
	"net/http"
)

func main() {
	// create a multiplexer for first page.
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveFirstPage)

	// create yet another multiplexer for creating zettels.
	mux.HandleFunc("/zettel/create", zettelCreate)

	// create yet another multiplexer for viewing zettels.
	mux.HandleFunc("/zettel/view", zettelView)

	log.Println("Starting the server at port 8082")
	err := http.ListenAndServe("localhost:8082", mux)
	log.Fatal(err)
}
