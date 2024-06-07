package main

import (
	"ascii-art-web/core"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, templ string, p *core.Page) {
	t, _ := template.ParseFiles(templ + ".html")
	t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	p, _ := core.LoadPage("test")
	renderTemplate(w, "index", p)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
