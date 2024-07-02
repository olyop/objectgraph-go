package parser

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/vektah/gqlparser/ast"
)

func readSchema(schemaFs fs.FS) ([]*ast.Source, error) {
	sourceFiles := make([]*ast.Source, 0)

	err := fs.WalkDir(schemaFs, ".", walkSourceFile(schemaFs, &sourceFiles))
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", "graphql/schema", err)
	}

	return sourceFiles, err
}

var validExtensions = map[string]struct{}{
	".graphql":  {},
	".graphqls": {},
	".gql":      {},
	".gqls":     {},
}

func walkSourceFile(schemaFs fs.FS, sourceFiles *[]*ast.Source) fs.WalkDirFunc {
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

		source := &ast.Source{
			Name:  path,
			Input: string(contents),
		}

		*sourceFiles = append(*sourceFiles, source)

		return nil
	}
}
