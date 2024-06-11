package files

import (
	"io"
	"io/fs"
	"log"
	"strings"
)

func processFolderMap(fs fs.FS, base []string) map[string][]string {
	folderMap := make(map[string][]string)

	for _, item := range base {
		file, err := fs.Open(item)
		if err != nil {
			log.Default().Fatal(err)
		}

		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			log.Default().Fatal(err)
		}

		contents := string(data)

		folderMap[item] = processFile(contents)
	}

	return folderMap
}

func processFileTabSeperated(contents string) [][]string {
	file := processFile(contents)

	fileValues := make([][]string, len(file))
	for i, line := range file {
		fileValues[i] = strings.Split(line, "\t")
	}

	return fileValues
}

func processFile(contents string) []string {
	values := strings.Split(contents, "\n")

	trimmedValues := make([]string, len(values))
	for i, value := range values {
		trimmedValues[i] = strings.TrimSpace(value)
	}

	return trimmedValues
}
