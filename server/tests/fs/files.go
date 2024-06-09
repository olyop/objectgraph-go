package fs

import (
	"log"
	"os"
	"path/filepath"
)

const (
	testFilesFolder = "fs"
)

func ImportTestFile(path string) string {
	file, err := os.ReadFile(filepath.Join(testFilesFolder, path))
	if err != nil {
		log.Default().Fatalf("fs: error %v", err)
	}

	content := string(file)

	return content
}
