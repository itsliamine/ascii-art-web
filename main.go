package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii", handlers.PostHandler)
	http.HandleFunc("/404", handlers.NotFoundHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
