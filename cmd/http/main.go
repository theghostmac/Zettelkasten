package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application holds all dependencies for Zettelkasten. Injecting first-page dependencies into main().
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a new cli flag with 'addr' to host a port.
	addr := flag.String("addr", ":8082", "HTTP network address")
	// Parse the flag to the CLI flag.
	flag.Parse()

	// Create a logger for INFO to Stdout.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for ERROR to Stderr.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	// create a multiplexer for first page.
	mux := http.NewServeMux()

	// create a fileServer and register it using mux.Handle() as a handler for all URL paths in /static/.
	fileServer := http.FileServer(http.Dir("./assets/static/"))
	// also strip the /static prefix away before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.serveFirstPage)
	mux.HandleFunc("/zettel/create", app.zettelCreate)
	// create yet another multiplexer for viewing zettels.
	mux.HandleFunc("/zettel/view", app.zettelView)

	serverProp := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting the server at port %s", *addr)
	err := serverProp.ListenAndServe()
	errorLog.Fatal(err)
}
