package main

import (
	"GustavoDelfim/graphql-go-example/middleware"
	"GustavoDelfim/graphql-go-example/resolver"
	"GustavoDelfim/graphql-go-example/schema"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Running"))
	})

	Schema := schema.GetSchema(&resolver.RootResolver{})

	router.Group(func(router chi.Router) {
		router.Use(middleware.AuthMiddleware)
		router.Handle("/graphql", &relay.Handler{Schema: Schema})
	})

	fmt.Println("Running in port ", 8080)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
