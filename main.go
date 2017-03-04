package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/itsshompi/kuubit-backend/handlers"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
	port        = ":8080"
	serverCert  = "server.crt"
	serverKey   = "server.key"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{name: Kuubit API, version: 1.0.1}"))
}

func main() {
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/check", handler.AuthHandler)
	fmt.Println("Server is running in https://localhost" + port)
	err := http.ListenAndServeTLS(port, serverCert, serverKey, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
