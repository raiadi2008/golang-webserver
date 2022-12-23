package main

import (
	"log"
	"net/http"

	"github.com/raiadi2008/golang-webserver/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("App started on port http://127.0.0.1:%s", port)
	_ = http.ListenAndServe(port, nil)
}
