package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ashwinp15/audio-directory/database"
	"github.com/ashwinp15/audio-directory/graph"
	"github.com/ashwinp15/audio-directory/graph/generated"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	defer database.PGclient.Close()

	router := chi.NewRouter()
	//router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// GUI testing endpoint
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// Single endpoint for all queries and mutations
	router.Handle("/query", srv)
	//http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
