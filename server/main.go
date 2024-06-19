package main

import (
	"embed"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/olyop/graphqlops-go/database"
	"github.com/olyop/graphqlops-go/graphql/resolvers"
	"github.com/olyop/graphqlops-go/graphql/retrievers"
	"github.com/olyop/graphqlops-go/graphqlops"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

//go:embed graphql/schema/*.graphqls
var schemaFs embed.FS

func main() {
	env := loadEnv()

	database.Connect()
	defer database.Close()

	config := &graphqlops.Configuration{
		Schema:     schemaFs,
		Resolvers:  &resolvers.Resolver{},
		Retrievers: &retrievers.Retrievers{},
		CacheRedis: &redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		CachePrefix: os.Getenv("REDIS_PREFIX"),
		CacheDurations: graphqlops.CacheDurationMap{
			"catalog": 2 * time.Minute,
		},
	}

	engine, err := graphqlops.NewEngine(config)
	if err != nil {
		panic(err)
	}

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
	r.Use(corsHandler())

	r.Handle("/graphql", graphqlops.Handler{Engine: engine})

	if env == "development" {
		err := http.ListenAndServeTLS(":8080", os.Getenv("TLS_CERT_PATH"), os.Getenv("TLS_KEY_PATH"), r)
		if err != nil {
			panic(err)
		}
	} else {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			panic(err)
		}
	}
}
