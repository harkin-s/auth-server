package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func getRoutes() *httprouter.Router {

	router := httprouter.New()

	router.GET("/", index)

	fmt.Println("Server strated")

	return router
}
