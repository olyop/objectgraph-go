package queries

import (
	"os"
	"path"
)

var SelectClassificationByIDQuery = readSQLFile("select-classification-by-id")
var SelectBrandByIDQuery = readSQLFile("select-brand-by-id")
var SelectCategoriesByProductID = readSQLFile("select-categories-by-product-id")
var SelectProductByID = readSQLFile("select-product-by-id")
var SelectProductsQuery = readSQLFile("select-products")

func readSQLFile(filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := path.Join(wd, "database", "queries", filename+".sql")

	contents, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(contents)
}
