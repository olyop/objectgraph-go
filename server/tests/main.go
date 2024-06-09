package main

import (
	"log"

	"github.com/olyop/graphql-go/tests/fs"
	"github.com/olyop/graphql-go/tests/graphql"
	"github.com/olyop/graphql-go/tests/utils"
)

func main() {
	log.SetFlags(log.Lshortfile)

	graphql.InitializeClient()

	runGraphQlTests()
}

func runGraphQlTests() {
	testGetProduct()
	testGetProductsLoad()
}

func testGetProduct() {
	file := fs.ImportTestFile("get-product.graphql")

	graphql.RunQuery(file)
}

func testGetProductsLoad() {
	file := fs.ImportTestFile("get-products.graphql")

	graphql.RunQuery(file)

	fn := func() {
		graphql.RunQuery(file)
	}

	utils.Iterate(fn, utils.IterateOptions{
		Iterations: 100,
		Pace:       100,
	})
}
