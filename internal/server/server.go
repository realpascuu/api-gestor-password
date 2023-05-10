package server

import (
	"crypto/tls"
	"gestorpasswordapi/certs"
	v1 "gestorpasswordapi/internal/server/v1"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

func (serv *Server) Close() error {
	serv.server.Close()
	return nil
}

func (serv *Server) Start() {
	certContents, err := serv.readKeyFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}

	keyContents, err := serv.readKeyFile("server.key")
	if err != nil {
		log.Fatal(err)
	}

	serv.addTLSConfig(certContents, keyContents)

	log.Printf("Server running on https://%s", serv.server.Addr)
	err = serv.server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

func (serv *Server) addTLSConfig(certContents []byte, keyContents []byte) error {
	cert, err := tls.X509KeyPair(certContents, keyContents)
	if err != nil {
		return err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	serv.server.TLSConfig = tlsConfig
	return nil
}

func (serv *Server) readKeyFile(keyPath string) ([]byte, error) {
	certsFS := &certs.Certs
	fileFS, err := certsFS.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	return fileFS, nil
}

func New(hostname, port string) (*Server, error) {
	r := chi.NewRouter()

	// * show logs of the request to the server
	r.Use(middleware.Logger)
	// * recovers the app when a panic occurs
	r.Use(middleware.Recoverer)

	// * API routes version 1
	r.Mount("/api/v1", v1.New())

	serv := &http.Server{
		Addr:         hostname + ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server := Server{server: serv}
	return &server, nil
}
