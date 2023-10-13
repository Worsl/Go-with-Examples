package handlers

import (
	"fmt"
	"net/http"

	"github.com/worsl/Go-Workspace/webapp/pkg/render"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

func Headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
