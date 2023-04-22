package main

import (
	"log"
	"net/http"
)

func holaMundo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	var str = "hola mundo!"
	writer.Write([]byte(str))
}

/* HTTPS PROTOCOL */
func main() {
	http.HandleFunc("/home", holaMundo)
	err := http.ListenAndServeTLS("localhost:443", "certs/server.crt", "certs/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
HTTP PROTOCOL
func main() {
	log.Println("Try to listen on port :8080")
	http.HandleFunc("/home", holaMundo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	} else {
		log.Println("Listening on port :8080")
	}

}
*/
