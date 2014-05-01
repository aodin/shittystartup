package shittystartup

import (
	"html/template"
	"net/http"
)

type Server struct {
	Output string
}

type Attrs struct {
	StaticURL string
}

var tmpl = template.Must(template.ParseFiles("./shittystartup/templates/landing.html"))

func (s *Server) WriteResponse(w http.ResponseWriter, r *http.Request) {
	attrs := Attrs{StaticURL: "/static/"}
	tmpl.Execute(w, attrs)
}

func NewServer(output string) *Server {
	return &Server{output}
}
