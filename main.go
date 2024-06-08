package main

import (
	core "ascii-art-web/core"
	"html/template"
	"log"
	"net/http"
)

// func main() {
// 	fmt.Println(asciiart.GetAscii("hello", "standard"))
// }

func renderTemplate(w http.ResponseWriter, p *core.Page) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	p, err := core.LoadPage("test")
	if err != nil {
		p = &core.Page{Title: "test"}
	}
	renderTemplate(w, p)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	ret := template.HTMLEscapeString("Hello\nworld")
	p := &core.Page{Title: "test", Body: ret}
	core.Save(p)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii", postHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
