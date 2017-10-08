package main

import (
	database "backend/employee/database"
	/*common "backend/employee/common"
	router "backend/employee/router"
	"log"
	"net/http"*/
)

func main() {
	database.Init()
	/*common.InitConfig()

	router := router.InitRoutes()

	server := &http.Server {
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Println("Listening...")
	server.ListenAndServe()*/
}
