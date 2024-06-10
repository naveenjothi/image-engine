package main

import (
	"log"
	"myapp/internal/app/middleware"
	"myapp/internal/config"
	"myapp/internal/router"
	"net/http"
	"os"
)

func main(){
	config.LoadConfig()

	port := os.Getenv("APP_PORT")
    if port == "" {
        log.Fatal("APP_PORT environment variable not set")
    }
	
    r := router.NewRouter()
	log.Printf("Starting server on :%s\n", port)
	r.Use(middleware.LoggingMiddleware)
    log.Fatal(http.ListenAndServe(":"+port, r))
}