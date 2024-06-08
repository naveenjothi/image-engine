package router

import (
	"myapp/internal/app/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
    r.HandleFunc("/users/create", handler.CreateUser).Methods("POST")
    return r
}