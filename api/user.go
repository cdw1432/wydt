package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wydt/data"

	"github.com/gorilla/mux"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	user, err := data.GetUserByHandle("test")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func RegisterUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/api/user").Subrouter()
	userRouter.HandleFunc("/me", GetUserInfo).Methods("GET")
}
