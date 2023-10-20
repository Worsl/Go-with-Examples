package main

import (
	"log"
	"net/http"
	"github.com/worsl/Go-Workspace/webapp/pkg/config"
	"github.com/worsl/Go-Workspace/webapp/pkg/handlers"
	"github.com/worsl/Go-Workspace/webapp/pkg/render"
)

func main() {

	tc,err := render.CreateTemplateCache() // the application cache is created once when the server starts
	if err!= nil{
		log.Fatal("cannot create template cache")
	}

	var app config.AppConfig
	render.NewTemplates(&app)

	app.TemplateCache = tc
	app.UseCache = false // should we allow program to use existing cache

	var repo *handlers.Repository = handlers.NewRepo(&app) 
	handlers.NewHandlers(repo)

	

	http.HandleFunc("/", handlers.Repo.Hello) // or http.HandleFunc("/", repo.Hello)
	http.HandleFunc("/headers", repo.Headers) // or http.HandleFunc("/headers", Render.Repo.Headers)
	http.HandleFunc("/about",handlers.Repo.About)

	http.ListenAndServe(":8080", nil)
}
