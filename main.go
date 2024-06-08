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

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	t, _ := template.ParseFiles("templates/index.html")
	data := PageData{}
	f, _ := os.ReadFile("content.txt")
	fileContent := string(f)
	data.Lines = strings.Split(fileContent, "\n")
	t.Execute(w, data)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	style := r.FormValue("banner")
	output := asciiart.GetAscii(input, style)
	core.Save(output)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func main() {
	http.HandleFunc("/home", handler)
	http.HandleFunc("/ascii", postHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
