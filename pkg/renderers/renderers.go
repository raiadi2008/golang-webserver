package renderers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tc = make(map[string]*template.Template)

func RenderTemplatesTest(w http.ResponseWriter, tmpl string) {
	parsedFiles, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		log.Println("error occured while parsing files", err)
	}
	parsedFiles.Execute(w, "some data")

}

func RenderTemplates(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if template exists in cache
	_, isPresent := tc[t]

	if isPresent {
		// the template is present in the cache so we can return
		log.Println("Template is present inside the cache")
	} else {
		log.Println("Create template and adding it to cache")
		err = createTemplate(t)
		if err != nil {
			log.Println("Error occured while creating template", err)
		}

	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error occured during writing resposne from cahce", err)
	}
}

func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl

	return nil
}
