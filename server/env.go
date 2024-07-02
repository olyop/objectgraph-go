package main

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

func loadEnv() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root := path.Dir(wd)
	dotEnvPath := path.Join(root, ".env")

	err = godotenv.Load(dotEnvPath)
	if err != nil {
		panic(err)
	}

	env := os.Getenv("GO_ENV")
	if env == "" {
		panic("GO_ENV is not set")
	}

	return env
}
