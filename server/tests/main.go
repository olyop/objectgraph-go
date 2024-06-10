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
	file1 := fs.ImportTestFile("get-product-1.graphql")
	file2 := fs.ImportTestFile("get-product-2.graphql")
	file3 := fs.ImportTestFile("get-product-3.graphql")
	file4 := fs.ImportTestFile("get-product-4.graphql")
	file5 := fs.ImportTestFile("get-product-5.graphql")
	file6 := fs.ImportTestFile("get-product-6.graphql")

	graphql.RunQuery(file1)
	graphql.RunQuery(file2)
	graphql.RunQuery(file3)
	graphql.RunQuery(file4)
	graphql.RunQuery(file5)
	graphql.RunQuery(file6)
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
