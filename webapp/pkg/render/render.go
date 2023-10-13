package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	// the first parameter takes in the extension of the base, while the 2nd paramter takes in the base itself
	// this may not be optimal because render and reading from disks takes up a lot of space, why not use a cache?
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error found", err)
	}
}

var pagemap = make(map[string]*template.Template)

func RenderTemplatesNew(w http.ResponseWriter, tmpl string) {
	parsedTemplate, exists := pagemap[tmpl]

	if !exists {
		parsedTemplate, _ = template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
		pagemap[tmpl] = parsedTemplate

	} else {
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error found", err)
		}
	}

}
