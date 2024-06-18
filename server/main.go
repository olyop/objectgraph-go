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
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/retrievers"
	"github.com/olyop/graphql-go/server/graphqlops"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

//go:embed graphql/schema/*.graphqls
var schemaFs embed.FS

func main() {
	env := loadEnv()

	database.Connect()
	defer database.Close()

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

	config := &graphqlops.Configuration{
		Schema:     schemaFs,
		Resolvers:  &resolvers.Resolver{},
		Retrievers: &retrievers.Retrievers{},
		Cache: &graphqlops.CacheConfiguration{
			Prefix: os.Getenv("REDIS_PREFIX"),
			Durations: graphqlops.CacheDurationMap{
				"catalog": 2 * time.Minute,
			},
			Redis: &redis.Options{
				Addr:     os.Getenv("REDIS_URL"),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       0,
			},
		},
	}

	engine, err := graphqlops.NewEngine(config)
	if err != nil {
		log.Fatal(err)
	}

	defer engine.Close()

	r.Handle("/graphql", graphqlops.Handler{Engine: engine})

	if env == "development" {
		err := http.ListenAndServeTLS(":8080", os.Getenv("TLS_CERT_PATH"), os.Getenv("TLS_KEY_PATH"), r)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Fatal(err)
		}
	}
}
