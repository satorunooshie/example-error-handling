package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/satorunooshie/example-error-handling/handler"
)

func main() {
	if err := http.ListenAndServe(":8080", router()); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}
}

func router() http.Handler {
	mux := http.NewServeMux()
	userHandler := handler.User{}
	mux.HandleFunc("GET /user/{id}", userHandler.Get)
	mux.HandleFunc("POST /user", userHandler.Create)
	mux.HandleFunc("DELETE /user/{id}", userHandler.Delete)
	return mux
}
