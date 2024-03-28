package main

import (
	"net/http"

	"github.com/satorunooshie/example-error-handling/handler"
)

func main() {
	mux := http.NewServeMux()
	userHandler := handler.User{}
	mux.HandleFunc("GET /user/{id}", userHandler.Get)
	mux.HandleFunc("POST /user", userHandler.Create)
	http.ListenAndServe(":8080", mux)
}
