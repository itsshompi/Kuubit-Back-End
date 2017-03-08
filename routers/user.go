package routers

import (
	"github.com/gorilla/mux"
	"github.com/itsshompi/kuubit-backend/controllers"
)

//SetUserRoutes is
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/signup", controllers.RegisterController).Methods("POST")
	router.HandleFunc("/auth/signin", controllers.LoginController).Methods("POST")
	return router
}
