package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql"
	"github.com/olyop/graphql-go/server/graphql/resolvers"

	_ "github.com/lib/pq"
)

//go:embed graphql/schema
var schemaFs embed.FS

func main() {
	loadEnv()

	database.Connect()
	defer database.Close()

	graphql.Initialize()

	engine.Initialize()
	defer engine.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentEncoding("gzip", "deflate"))
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.Compress(5, "application/json", "application/problem+json"))
	r.Use(middleware.NoCache)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(time.Second * 30))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(corsHandler())
	r.Post("/graphql", engine.CreateHandler(schemaFs, &resolvers.Resolver{}))

	if os.Getenv("GO_ENV") == "development" {
		log.Fatal(http.ListenAndServeTLS(":8080", os.Getenv("TLS_CERT_PATH"), os.Getenv("TLS_KEY_PATH"), r))
	} else {
		log.Fatal(http.ListenAndServe(":8080", r))
	}
}
