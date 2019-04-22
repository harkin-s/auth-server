package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	connection := getDbConnection()

	// Only yo use the object for now
	fmt.Println("Db connection %s", &connection)

	fmt.Fprint(w, "All looks good")

}
