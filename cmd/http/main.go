package main

import (
	"log"
	"net/http"
)

func main() {
	// create a multiplexer for first page.
	mux := http.NewServeMux()

	// create a fileServer and register it using mux.Handle() as a handler for all URL paths in /static/.
	fileServer := http.FileServer(http.Dir("./assets/static/"))
	// also strip the /static prefix away before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", serveFirstPage)

	// create yet another multiplexer for creating zettels.
	mux.HandleFunc("/zettel/create", zettelCreate)

	// create yet another multiplexer for viewing zettels.
	mux.HandleFunc("/zettel/view", zettelView)

	log.Println("Starting the server at port 8082")
	err := http.ListenAndServe("localhost:8082", mux)
	log.Fatal(err)
}
