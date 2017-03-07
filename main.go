package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/itsshompi/kuubit-backend/routers"
)

const (
	privKeyPath = "keys/privKey.key"
	pubKeyPath  = "keys/pubKey.pem"
	port        = ":9999"
)

func main() {
	//common.StartUp() - Replaced with init method
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	fmt.Println("Server is running in https://localhost" + port)
	err := http.ListenAndServeTLS(port, pubKeyPath, privKeyPath, n)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
