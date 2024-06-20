package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (page * Page) save() error {
	filename := page.Title + ".txt" //Remember to refer to page.Title 
	return os.WriteFile(filename, page.Body, 0600)
}

func load_page(title string) (*Page, error){
	filename := title + ".txt"
	body, error := os.ReadFile(filename)

	if error != nil {
		return nil, error
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "Cool", Body: []byte("This is a cool page")}
	p1.save()
	p2, _ := load_page("Cool")
	fmt.Println(string(p2.Body))
}
