package handlers

import (
	"fmt"
	"net/http"

	"github.com/worsl/Go-Workspace/webapp/pkg/config"
	"github.com/worsl/Go-Workspace/webapp/pkg/models"
	"github.com/worsl/Go-Workspace/webapp/pkg/render"
)


var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(appFromMain *config.AppConfig) (*Repository) {
	return &Repository{
		App:appFromMain, 
	}
}

// NewHandlers sets the repository for the Handlers
func NewHandlers(repoFromMain *Repository){
	Repo = repoFromMain
}

// a function reciever associates itself with the Repository struct, allowing it to have access to its fields.
func (m *Repository) Hello(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, "home.page.html",&models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, req *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
