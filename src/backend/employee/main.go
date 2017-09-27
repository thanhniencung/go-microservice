package main

import (
	common "backend/employee/common"
	router "backend/employee/router"
	"log"
	"net/http"
)

func main() {
	common.InitConfig()

	// Get the mux router object
	router := router.InitRoutes()

	server := &http.Server {
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
