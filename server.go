package shittystartup

import (
	"net/http"
)

type Server struct {
	Output string
}

func (s *Server) WriteResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.Output))
}

func NewServer(output string) *Server {
	return &Server{output}
}
