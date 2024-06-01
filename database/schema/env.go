package main

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

func loadEnv() {
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
}
