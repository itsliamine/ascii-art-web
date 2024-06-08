package core

import "os"

type Page struct {
	Title string
	Body  string
}

func Save(p *Page) error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, []byte(p.Body), 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: string(body)}, nil
}
