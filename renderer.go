package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplates(w http.ResponseWriter, tmpl string) {
	parsedFiles, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Println("error occured while parsing files", err)
	}
	parsedFiles.Execute(w, "some data")
}
