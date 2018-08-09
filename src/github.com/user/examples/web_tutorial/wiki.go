package main

import (
	"io/ioutil"
)

//Page struct
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	data, err := ioutil.ReadFile(title + ".txt")

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: data}, nil
}

// func main() {
// 	page := Page{
// 		Title: "test",
// 		Body:  []byte("This is a sample test"),
// 	}

// 	page.save()
// 	p2, _ := loadPage(page.Title)
// 	fmt.Println(string(p2.Body))
// }
