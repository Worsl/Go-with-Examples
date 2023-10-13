package handlers

import (
	"fmt"
	"net/http"
	"github.com/worsl/Go-Workspace/webapp/pkg/render"
)



func Hello(w http.ResponseWriter, req *http.Request) {

	render.RenderTemplates(w,"home")
}

func Headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}


