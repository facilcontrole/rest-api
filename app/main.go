package main

import (
	"log"
	"net/http"

	"github.com/facilcontrole/rest-api/app/middleware"
	"github.com/facilcontrole/rest-api/app/routes"
)

func main() {

	port := "1972"
	http.Handle("/", middleware.Cors(routes.Routes()))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
