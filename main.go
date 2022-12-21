package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const PORT = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Page")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Printf("error occured while parsing html: %s", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(PORT, nil)
}
