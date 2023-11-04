package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/worsl/Go-Workspace/webapp/pkg/config"
	"github.com/worsl/Go-Workspace/webapp/pkg/models"
)

var pagemap = make(map[string]*template.Template)
var app *config.AppConfig

// NewTemplates sets the config for the configApplication
func NewTemplates(appFromMain *config.AppConfig) {
	app = appFromMain // successfully linked both of them together
}


func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	td.StringMap["bye"] = "bye"
	return td
}

// RenderTemplates have an advantage in that we can have multiple base layouts and it will be handled accordingly, with ParseGlob later
func RenderTemplate(w http.ResponseWriter, tmpl string , td *models.TemplateData) {

	var pageCache map[string]*template.Template
	// if the system configuration allows us to use the existing cache, 
	if app.UseCache{
		// get the template cache from app config, instead of creating it everytime
		pageCache = app.TemplateCache
	}else{
		pageCache,_ = CreateTemplateCache()
	}



	// get request template from cache
	t, exists := pageCache[tmpl]
	if !exists {
		log.Fatal("page not found in cache")
	}

	td = AddDefaultData(td)

	// render the template
	err := t.Execute(w, td)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html first from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	// range through all the files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)                    // returns the last element of the path(filename) instead of entire path
		ts, err := template.New(name).ParseFiles(page) // recall that page is the directory

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}


func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, exists := pagemap[tmpl]

	if !exists {
		fmt.Printf("page {%s} not found in cache, parsing it and adding to cache\n", tmpl)
		parsedTemplate, _ = template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html") // what if multiple bases?
		pagemap[tmpl] = parsedTemplate

	}

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error found", err)
	}

}
