package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wydt/data"

	"github.com/gorilla/mux"
)

func GetObjectsInfo(w http.ResponseWriter, r *http.Request) {
	objects, err := data.GetObjectByHandle("test")

	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(objects)
}

func RegisterObjectsRoutes(r *mux.Router) {
	objRouter := r.PathPrefix("/api/objects").Subrouter()
	objRouter.HandleFunc("/show", GetObjectsInfo).Methods("GET")
}
