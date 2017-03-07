package routers

import (
	"github.com/gorilla/mux"
)

//InitRoutes is ...
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	// Routes for the User entity
	router = SetUserRoutes(router)

	return router
}
