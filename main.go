package main

import (
	"log"
	"net/http"

	handler "github.com/itsshompi/kuubit-backend/handlers"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{name: Kuubit API, version: 1.0.1}"))
}

func main() {
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/login", handler.LoginHandler)
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
