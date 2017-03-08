package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/itsshompi/kuubit-backend/common"
	"github.com/itsshompi/kuubit-backend/routers"
)

const (
	privKeyPath = "keys/privKey.key"
	pubKeyPath  = "keys/pubKey.pem"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	fmt.Println("Server is running in https://" + common.AppConfig.Server)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	err := server.ListenAndServeTLS(pubKeyPath, privKeyPath)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
