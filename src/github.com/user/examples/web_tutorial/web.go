package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hiii , I am at %s", r.URL.Path[1:])

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	// log.Print(title)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, html string, p *Page) {
	t, err := template.ParseFiles(html + ".html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err2 := t.Execute(w, p)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func main() {
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
