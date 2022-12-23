package handlers

import (
	"net/http"

	"github.com/raiadi2008/golang-webserver/pkg/renderers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplates(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplates(w, "about.html")
}
