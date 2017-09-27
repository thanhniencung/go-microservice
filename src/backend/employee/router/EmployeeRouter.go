package router

import (
	"backend/employee/controller"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUsersRouters(router)
	return router
}

func SetUsersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controller.AddNewEmployee).Methods("POST")
	return router
}
