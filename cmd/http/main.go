package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// HandleError receives and handles all errors in this project.
func HandleError(err error) {
	if err != nil {
		fmt.Println("Error occurred: ", err)
	}
}

// zettelView will view a given zettel via id query from the user.
func zettelView(writer http.ResponseWriter, reader *http.Request) {
	// receive the id value and convert to an integer. If it's not an integer or
	// if it's not non-zero positive integer, return a 404 page not found response.
	id, err := strconv.Atoi(reader.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, reader)
		return
	}
	_, err = fmt.Fprintf(writer, "Display a specific zettel with ID %d", id)
	HandleError(err)
}

// zettelCreate gives you the power to create new zettel on your Zettelkasten.
func zettelCreate(writer http.ResponseWriter, reader *http.Request) {
	// if the method is not a POST request, return an error message.
	if reader.Method != "POST" {
		// use Header().Set() method to add a "Allow", "POST" map to the header.
		writer.Header().Set("Allow", http.MethodPost) // replaced "POST" with net/http constant: http.MethodPost.
		/* REPLACE ENTIRE:
			writer.WriteHeader(405)
			_, err := writer.Write([]byte("That HTTP method is not supported on this handler function."))
			HandleError(err)
		WITH:
		*/
		http.Error(writer, "Method is not supported", http.StatusMethodNotAllowed) // http.Error is a shortcut.
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