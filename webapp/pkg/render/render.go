package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


var pagemap = make(map[string]*template.Template)

func RenderTemplatesOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, exists := pagemap[tmpl]

	if !exists {
		fmt.Printf("page {%s} not found in cache, parsing it and adding to cache\n",tmpl)
		parsedTemplate, _ = template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html") // what if multiple bases?
		pagemap[tmpl] = parsedTemplate

	}


	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error found", err)
	}

}


// RenderTemplatesNew have an advantage in that we can have multiple base layouts and it will be handled accordingly, with ParseGlob later
func RenderTemplates(w http.ResponseWriter, tmpl string){

	// create a template cache (only if it first request?)
	pageCache, err := createTemplateCache()
	if err!= nil{
		log.Fatal(err)
	}

	// get request template from cache
	t,exists := pageCache[tmpl]
	if !exists{
		log.Fatal(err)
	}

	// render the template
	err = t.Execute(w,nil)

	if err!= nil{
		fmt.Println("error parsing template:", err)
	}


}


func createTemplateCache() (map[string]*template.Template,error){

	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html first from ./templates
	pages,err := filepath.Glob("./templates/*.page.html")

	if err!= nil{
		return myCache,err
	}

	// range through all the files ending with *.page.html
	for _,page:= range(pages){
		name := filepath.Base(page) // returns the last element of the path(filename) instead of entire path
		ts,err := template.New(name).ParseFiles(page) // recall that page is the directory
 
		if err!= nil{
			return myCache,err
		}

		matches,err := filepath.Glob("./templates/*.layout.html")
		if err!=nil{
			return myCache,err
		}

		if len(matches) > 0{
			ts,err = ts.ParseGlob("./templates/*.layout.html")
			if err!=nil{
				return myCache,err
			}
		}

		myCache[name] = ts
	}
	return myCache,nil

	
}