package main

import (
	"fmt"
	"log"
	"net/http"
)

// HandleError receives and handles all errors in this project.
func HandleError(err error) {
	if err != nil {
		fmt.Println("Error occurred: ", err)
	}
}

// zettelView defines a handler function for viewing your zettels.
func zettelView(writer http.ResponseWriter, reader *http.Request) {
	_, err := writer.Write([]byte("Hello from inside the simple server..."))
	HandleError(err)
}

// zettelCreate gives you the power to create new zettel on your Zettelkasten.
func zettelCreate(writer http.ResponseWriter, reader *http.Request) {
	_, err := writer.Write([]byte("Create a new zettel here..."))
	HandleError(err)
}

func serveFirstPage(writer http.ResponseWriter, reader *http.Request) {
	_, err := writer.Write([]byte("Displaying a second page for Zettelkasten..."))
	HandleError(err)
}

func main() {
	// create a multiplexer for first page.
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveFirstPage)

	// create yet another multiplexer for creating zettels.
	mux.HandleFunc("/zettel/create", zettelCreate)

	// create yet another multiplexer for viewing zettels.
	mux.HandleFunc("/zettel/view", zettelView)

	log.Println("Starting the server at port 8082")
	err := http.ListenAndServe("localhost:8082", nil)
	log.Fatal(err)
}
