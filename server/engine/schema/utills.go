package schema

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func readSchema(schemaFs fs.FS) (schema string, err error) {
	sourceFiles := make(map[string]*sourceFile)

	err = fs.WalkDir(schemaFs, ".", walkSourceFile(schemaFs, sourceFiles))
	if err != nil {
		return schema, fmt.Errorf("failed to walk directory %s: %w", "graphql/schema", err)
	}

	for _, sourceFile := range sourceFiles {
		schema += sourceFile.contents
	}

	return schema, err
}

func walkSourceFile(schemaFs fs.FS, sourceFiles map[string]*sourceFile) fs.WalkDirFunc {
	var validExtensions = map[string]struct{}{
		".graphql":  {},
		".gql":      {},
		".graphqls": {},
		".gqls":     {},
	}

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

		file, err := schemaFs.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		contents, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		sourceFiles[path] = &sourceFile{
			path:     path,
			contents: string(contents),
		}

		return nil
	}
}
