package shittystartup

import (
	"html/template"
	"net/http"
)

type Server struct {
	tmpl *template.Template
}

type Attrs struct {
	StaticURL string
}

func (s *Server) WriteResponse(w http.ResponseWriter, r *http.Request) {
	attrs := Attrs{StaticURL: "/static/"}
	s.tmpl.Execute(w, attrs) // Execute returns an error that we're ignoring
}

func NewServer() (*Server, error) {
	// Create a pointer to a new, zero-initialized Server struct
	server := &Server{}

	// Since "tmpl" is lowercase it can only be accessed within the struct
	tmpl, err := template.ParseFiles("./shittystartup/templates/landing.html")

	// The standard error checking practice
	if err != nil {
		// Either return the "incomplete" server or nil, it's up to you
		// The user of the package should be checking for returned errors
		// before making use of the returned server struct.
		return server, err
	}

	// We don't use assignment (:=) because the variable "server.tmpl"
	// already exists. It was zero initialized in the server struct.
	server.tmpl = tmpl

	// If there is no error to return, return nil
	return server, nil
}
