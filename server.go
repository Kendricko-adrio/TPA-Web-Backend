package main

import (
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kendricko-adrio/gqlgen-todos/graph/generated"
	"github.com/kendricko-adrio/gqlgen-todos/graph/resolver"
)

const defaultPort = "2000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	//migration.SeedAll()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{} }))


	wrapped := middleware.CorsMiddleware(srv)
	wrapped2 := middleware.LoginMiddleware(wrapped)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", wrapped2)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
