package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tckthecreator/go-api-with-pgsql/configs"
	"github.com/tckthecreator/go-api-with-pgsql/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Post("/", handlers.Create)
	router.Get("/{id}", handlers.Get)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
