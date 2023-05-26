package main

import (
	"log"
	"net/http"
)

// serveFirstPage defines a handler function for serving the first page.
func serveFirstPage(writer http.ResponseWriter, reader *http.Request) {
	writer.Write([]byte("Hello from inside the simple server..."))
}

func main() {
	// create a multiplexor.
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveFirstPage)

	log.Println("Starting the server at port 8082")
	err := http.ListenAndServe(":8082", nil)
	log.Fatal(err)
}
