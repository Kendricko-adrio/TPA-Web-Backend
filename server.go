package main

import (
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/kendricko-adrio/gqlgen-todos/middleware"
	"log"
	"net/http"
	"os"
	"time"

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

	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{} }))

	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver.NewResolver(),
	}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})
	srv.Use(apollotracing.Tracer{})

	wrapped := middleware.CorsMiddleware(srv)
	wrapped2 := middleware.LoginMiddleware(wrapped)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", wrapped2)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
