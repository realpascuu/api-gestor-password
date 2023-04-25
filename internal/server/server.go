package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

func (serv *Server) Close() error {
	// TODO: add resource closure
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server running on https://%s", serv.server.Addr)
	err := serv.server.ListenAndServeTLS("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

func New(hostname, port string) (*Server, error) {
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         hostname + ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server := Server{server: serv}
	return &server, nil
}
