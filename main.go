package main

import (
	"net/http"
	"wydt/api"
	"wydt/data"

	"github.com/gorilla/mux"
)

func main() {
	data.DBInit()
	r := mux.NewRouter().StrictSlash(true)

	api.RegisterUserRoutes(r)
	api.RegisterObjectsRoutes(r)

	http.ListenAndServe("localhost:8080", r)
}
