package routers

import (
	"github.com/gorilla/mux"
	"github.com/itsshompi/kuubit-backend/controllers"
)

//SetHomeRoutes is
func SetHomeRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.HomeController).Methods("GET")
	return router
}
