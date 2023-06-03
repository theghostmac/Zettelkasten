package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new cli flag with 'addr' to host a port.
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parse the flag to the CLI flag.
	flag.Parse()
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

	log.Printf("Starting the server at port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
