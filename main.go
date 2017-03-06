package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/itsshompi/kuubit-backend/handlers"
)

const (
	privKeyPath = "keys/server.crt"
	pubKeyPath  = "keys/server.key"
	port        = ":8080"
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
	err := http.ListenAndServeTLS(port, privKeyPath, pubKeyPath, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
