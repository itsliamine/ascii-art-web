package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/ascii", handlers.PostHandler)
	mux.HandleFunc("/404", handlers.NotFoundHandler)

	// Custom 404 handler
	mux.HandleFunc("/notfound", handlers.NotFoundHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server started at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
