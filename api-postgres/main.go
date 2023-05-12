package main

import (
	"configs"
	"fmt"
	"handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	fmt.Println("Chegou no ListenAndServe", configs.GetDB())
	fmt.Println("API", configs.GetAPI())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}