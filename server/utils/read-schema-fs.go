package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/vektah/gqlparser/ast"
)

func CompileSchemaFS(schemaFs fs.FS) string {
	return compileSchemaFiles(readSchema(schemaFs))
}

func compileSchemaFiles(schemaFiles []*ast.Source) string {
	schema := ""

	for _, source := range schemaFiles {
		schema += source.Input
	}

	return schema

}

func readSchema(schemaFs fs.FS) []*ast.Source {
	sourceFiles := make([]*ast.Source, 0)

	err := fs.WalkDir(schemaFs, ".", walkSourceFile(schemaFs, &sourceFiles))
	if err != nil {
		panic(fmt.Errorf("failed to walk directory %s: %w", "graphql/schema", err))
	}

	return sourceFiles
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
