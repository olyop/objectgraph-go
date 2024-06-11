package fs

import (
	"log"
	"os"
	"path/filepath"
)

const (
	testFilesFolder = "fs"
)

var cache = make(map[string]string)

func ImportTestFile(path string) string {
	if content, ok := cache[path]; ok {
		return content
	}

	file, err := os.ReadFile(filepath.Join(testFilesFolder, path))
	if err != nil {
		log.Default().Fatalf("fs: error %v", err)
	}

	content := string(file)

	cache[path] = content

	return content
}
