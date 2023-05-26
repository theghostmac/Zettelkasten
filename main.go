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
	// if the method is not a POST request, return an error message.
	if reader.Method != "POST" {
		// use Header().Set() method to add a "Allow", "POST" map to the header.
		writer.Header().Set("Allow", "POST")
		writer.WriteHeader(405)
		_, err := writer.Write([]byte("That HTTP method is not supported on this handler function."))
		HandleError(err)
		// return from the function so the subsequent code is not executed.
		return
	}
	_, err := writer.Write([]byte("Create a new zettel here..."))
	HandleError(err)
}

// serveFirstPage serves the first page seen when a visitor visits Zettelkasten.
func serveFirstPage(writer http.ResponseWriter, reader *http.Request) {
	// if a requested URL path matches '/', return. If not, return http.NotFound().
	if reader.URL.Path != "/" {
		http.NotFound(writer, reader)
		return
	}

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
	err := http.ListenAndServe("localhost:8082", mux)
	log.Fatal(err)
}
