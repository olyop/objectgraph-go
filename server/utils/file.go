package utils

import (
	"os"
	"path"
)

func ReadFile(module string, filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var moduleFolder string
	if module == "database" {
		moduleFolder = path.Join("database", "queries")
	} else {
		moduleFolder = "schema"
	}

	filePath := path.Join(wd, moduleFolder, filename)

	contents, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(contents)
}
