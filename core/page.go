package core

import "os"

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}
