package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("App started on port http://127.0.0.1:%s", port)
	_ = http.ListenAndServe(port, nil)
}
