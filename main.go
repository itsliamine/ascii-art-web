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

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check correct GET method
	if r.Method != "GET" {
		log.Printf("Bad request %v on %v page\n", r.Method, r.URL.Path)
		badRequestHandler(w)
		return
	}

	if r.URL.Path != "/" {
		log.Printf("Tried to access unexistant route %v\n", r.URL.Path)
		notFoundHandler(w)
		return
	}

	// Parse the template for the home page
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing home template: %v", err)
		internalServerErrorHandler(w)
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
		internalServerErrorHandler(w)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	//Check correct POST method
	if r.Method != "POST" {
		log.Printf("Bad request %v on %v page\n", r.Method, r.URL.Path)
		badRequestHandler(w)
		return
	}

	// Parse form values
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing data form: %v", err)
		internalServerErrorHandler(w)
		return
	}

	// Retrieve input and banner style from the form
	input := r.FormValue("input")
	style := r.FormValue("banner")

	input = strings.Replace(input, "\r\n", "\n", -1)

	if style == "" {
		log.Printf("No banner provided: style: %s\n", style)
		internalServerErrorHandler(w)
		return
	}

	// Generate ASCII art based on input and style
	output := asciiart.GetAscii(input, style)

	// Save the generated ASCII art
	err = core.Save(output)
	if err != nil {
		log.Printf("Error saving output: %v", err)
		internalServerErrorHandler(w)
		return
	}

	// Redirect to the home page after successful processing
	http.Redirect(w, r, "/", http.StatusFound)
}

func notFoundHandler(w http.ResponseWriter) {
	// Send 404 code
	w.WriteHeader(http.StatusNotFound)

	// Parse the 404 template
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		log.Printf("Error executing 404 template: %v", err)
		internalServerErrorHandler(w)
		return
	}

	// Execute the 404 template
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing 404 template: %v", err)
		internalServerErrorHandler(w)
		return
	}
}

func badRequestHandler(w http.ResponseWriter) {
	// Send 400 code
	w.WriteHeader(http.StatusNotFound)

	// Parse the 400 template
	t, err := template.ParseFiles("templates/400.html")
	if err != nil {
		log.Printf("Error executing 400 template: %v", err)
		internalServerErrorHandler(w)
		return
	}

	// Execute the 400 template
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing 400 template: %v", err)
		internalServerErrorHandler(w)
		return
	}
}

func internalServerErrorHandler(w http.ResponseWriter) {
	// Send 500 code
	w.WriteHeader(http.StatusNotFound)

	// Parse the 400 template
	t, err := template.ParseFiles("templates/500.html")
	if err != nil {
		log.Printf("Error executing 500 template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the 400 template
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing 500 template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
