package main

import (
	"github.com/aodin/shittystartup/server"
	"log"
)

func main() {
	// It can't be called "server" because that's the package name!
	s, err := server.New()
	if err != nil {
		panic(err)
	}
	log.Println("Starting server...")
	s.ListenAndServe()
}
