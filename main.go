package main

import (
	asciiart "ascii-art-web/ascii-art"
	"ascii-art-web/core"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type PageData struct {
	Lines []string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", postHandler)
	http.HandleFunc("/404", notFoundHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}

	// Parse the template for the home page
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing home template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Read content from a file (e.g., content.txt) to display on the home page
	f, err := os.ReadFile("content.txt")
	if err != nil {
		log.Printf("Error reading content file: %v", err)
		// If the content file is not found or inaccessible, proceed without displaying content
	}
	fileContent := string(f)
	data := PageData{
		Lines: strings.Split(fileContent, "\n"),
	}

	// Execute the template with the provided data
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Error executing home template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Bad Request: Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Retrieve input and banner style from the form
	input := r.FormValue("input")
	style := r.FormValue("banner")

	// Generate ASCII art based on input and style
	output := asciiart.GetAscii(input, style)
	if len(output) == 0 {
		log.Printf("No output generated for input: %s, style: %s", input, style)
		http.Error(w, "Banner Not Found", http.StatusNotFound)
		return
	}

	// Save the generated ASCII art
	err = core.Save(output)
	if err != nil {
		log.Printf("Error saving output: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Redirect to the home page after successful processing
	http.Redirect(w, r, "/", http.StatusFound)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	// Parse the 404 template
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		log.Printf("Error parsing 404 template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the 404 template
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing 404 template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
