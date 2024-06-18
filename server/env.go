package main

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func loadEnv() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Default().Fatal("Error getting working directory")
	}

	root := path.Dir(wd)
	dotEnvPath := path.Join(root, ".env")

	err = godotenv.Load(dotEnvPath)
	if err != nil {
		log.Default().Fatal("Error loading .env file")
	}

	env := os.Getenv("GO_ENV")
	if env == "" {
		log.Default().Fatal("GO_ENV is not set")
	}

	return env
}
