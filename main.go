package main

import (
	"fmt"
	"log"
	"net/http"
)

// serveFirstPage defines a handler function for serving the first page.
func serveFirstPage(writer http.ResponseWriter, reader *http.Request) {
	_, err := writer.Write([]byte("Hello from inside the simple server..."))
	if err != nil {
		// Handle the error
		fmt.Println("Error occurred:", err)
	}
}

func main() {
	// create a multiplexer.
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveFirstPage)

	log.Println("Starting the server at port 8082")
	err := http.ListenAndServe(":8082", nil)
	log.Fatal(err)
}
