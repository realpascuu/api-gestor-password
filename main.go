package main

import (
	"os"
	"os/signal"

	"gestorpasswordapi/internal/data"
	"gestorpasswordapi/internal/server"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	serv, err := server.New(hostname, port)
	if err != nil {
		log.Fatal(err)
	}

	// connection to database
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	serv.Close()
	data.Close()
}
