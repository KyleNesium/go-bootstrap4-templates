package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("header.html", "footer.html", "main.html", "about.html"))

type Page struct {
	Title string
}

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "main", &Page{Title: "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "about", &Page{Title: "About"})
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":1337", nil)
}
