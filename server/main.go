package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

func main() {
	env := loadEnv()

	err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	engine, err := objectgraph.NewEngine(engineConfiguration())
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
	r.Handle("/graphql", engine)

	if env == "production" {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			panic(err)
		}
	} else {
		err := http.ListenAndServeTLS(":8080", os.Getenv("TLS_CERT_PATH"), os.Getenv("TLS_KEY_PATH"), r)
		if err != nil {
			panic(err)
		}
	}
}
