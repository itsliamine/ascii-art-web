package core

import "os"

type Page struct {
	Title string
	Body  []byte
}

func Save(lines []string) error {
	str := ""
	for _, line := range lines {
		str += line + "\n"
	}
	return os.WriteFile("content.txt", []byte(str), 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}
