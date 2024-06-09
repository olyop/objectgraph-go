package graphql

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var validExtensions = map[string]struct{}{
	".graphql":  {},
	".gql":      {},
	".graphqls": {},
	".gqls":     {},
}

func ReadSchema(path string) (schema string, err error) {
	sourceFiles := make(map[string]*SourceFile)

	err = filepath.WalkDir(path, walkSourceFile(sourceFiles))
	if err != nil {
		return schema, fmt.Errorf("failed to walk directory %s: %w", path, err)
	}

	for _, sourceFile := range sourceFiles {
		schema += sourceFile.Contents
	}

	return schema, err
}

func walkSourceFile(sourceFiles map[string]*SourceFile) fs.WalkDirFunc {
	return func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if _, ok := validExtensions[filepath.Ext(path)]; !ok {
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		sourceFiles[path] = &SourceFile{
			Path:     path,
			Contents: string(contents),
		}

		return nil
	}
}

// SourceFile represents a schema file
type SourceFile struct {
	Path     string
	Contents string
}
