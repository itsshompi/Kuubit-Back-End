package routers

import (
	"github.com/gorilla/mux"
	"github.com/itsshompi/kuubit-backend/controllers"
)

//SetUserRoutes is
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/signup", controllers.RegisterController).Methods("POST")
	router.HandleFunc("/login", controllers.LoginController).Methods("POST")
	router.HandleFunc("/home", controllers.HomeController).Methods("GET")
	return router
}
