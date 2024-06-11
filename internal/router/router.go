package router

import (
	"myapp/internal/app/handler"
	"myapp/internal/app/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    r.Use(middleware.LoggingMiddleware)
    // Public routes
    r.HandleFunc("/public/create-user", handler.CreateUser).Methods("POST")

    // protected routes
    protected := r.PathPrefix("/protected").Subrouter()
    protected.Use(middleware.JWTMiddleware)
    protected.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
    protected.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
    return r
}