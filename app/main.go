package main

import (
	"log"
	"net/http"

	"github.com/facilcontrole/rest-api/database/postgres"

	"github.com/facilcontrole/rest-api/app/middleware"
	"github.com/facilcontrole/rest-api/app/routes"
)

func main() {

	port := "1972"
	conn := postgres.App()
	http.Handle("/", middleware.Cors(routes.Routes(conn)))
	defer conn.Close()

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
