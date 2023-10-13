package main

import (
	"net/http"

	"github.com/worsl/Go-Workspace/webapp/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Hello)
	http.HandleFunc("/headers", handlers.Headers)
	http.HandleFunc("/about",handlers.About)

	http.ListenAndServe(":8080", nil)
}
