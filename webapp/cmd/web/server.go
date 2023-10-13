package main

import (
	"log"
	"net/http"

	"github.com/worsl/Go-Workspace/webapp/pkg/config"
	"github.com/worsl/Go-Workspace/webapp/pkg/handlers"
	"github.com/worsl/Go-Workspace/webapp/pkg/render"
)

func main() {
	var app config.AppConfig
	render.NewTemplates(&app)

	tc,err := render.CreateTemplateCache() // the application cache is created once when the server starts

	if err!= nil{
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Hello)
	http.HandleFunc("/headers", handlers.Headers)
	http.HandleFunc("/about",handlers.About)

	http.ListenAndServe(":8080", nil)
}
