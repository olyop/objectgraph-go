package parser

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"github.com/olyop/graphqlops-go/graphqlops/graphql"
)

func Exec(schemaFs fs.FS, resolver interface{}) (*graphql.Schema, error) {
	schemaString, err := readSchema(schemaFs)
	if err != nil {
		return nil, err
	}

	options := []graphql.SchemaOpt{
		graphql.MaxParallelism(runtime.GOMAXPROCS(0)),
		graphql.UseFieldResolvers(),
	}

	schema, err := graphql.ParseSchema(schemaString, resolver, options...)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

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

var validExtensions = map[string]struct{}{
	".graphql":  {},
	".gql":      {},
	".graphqls": {},
	".gqls":     {},
}

func walkSourceFile(schemaFs fs.FS, sourceFiles map[string]*sourceFile) fs.WalkDirFunc {
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

type sourceFile struct {
	path     string
	contents string
}
