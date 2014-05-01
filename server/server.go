package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type Server struct {
	tmpl   *template.Template
	config Config
}

// Our main HTTP response handler that will attached to the root URL.
func (s *Server) WriteResponse(w http.ResponseWriter, r *http.Request) {
	// Execute returns an error that we're currently ignoring
	s.tmpl.Execute(w, s.config)
}

// Wrap the http package's ListenAndServe
func (s *Server) ListenAndServe() error {
	// Build the server address using the config's Port number
	address := fmt.Sprintf(":%d", s.config.Port)
	return http.ListenAndServe(address, nil)
}

// The server constructor function
func New() (*Server, error) {
	// Create the new config
	config := ParseConfig()

	// Create a pointer to a new, zero-initialized Server struct
	server := &Server{config: config}

	// Since "tmpl" is lowercase it can only be accessed within the struct
	templatePath := filepath.Join(config.TemplateDir, "landing.html")
	tmpl, err := template.ParseFiles(templatePath)

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

	// Build the routes in the server constructor
	http.HandleFunc("/", server.WriteResponse)

	// Serve static files using the config's StaticURL
	http.Handle(
		config.StaticURL,
		http.StripPrefix(
			config.StaticURL,
			http.FileServer(http.Dir(config.StaticDir)),
		),
	)

	// If there is no error to return, return nil
	return server, nil
}
