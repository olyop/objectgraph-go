package main

import (
	"log"

	"github.com/olyop/objectgraph-go/tests/fs"
	"github.com/olyop/objectgraph-go/tests/graphql"
	"github.com/olyop/objectgraph-go/tests/utils"
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
	file1 := fs.ImportTestFile("get-product-1.graphql")
	file2 := fs.ImportTestFile("get-product-2.graphql")
	file3 := fs.ImportTestFile("get-product-3.graphql")
	file4 := fs.ImportTestFile("get-product-4.graphql")
	file5 := fs.ImportTestFile("get-product-5.graphql")
	file6 := fs.ImportTestFile("get-product-6.graphql")

	graphql.RunQuery(file)

	fn := func() {
		r1 := utils.RandomInt(1, 10)
		r2 := utils.RandomInt(3, 60)

		for i := 0; i < r1; i++ {
			graphql.RunQuery(file)
			utils.Sleep(r2)
			graphql.RunQuery(file1)
			utils.Sleep(r2 / 2)
			graphql.RunQuery(file2)
			utils.Sleep(r2 / 4)
			graphql.RunQuery(file3)
			graphql.RunQuery(file4)
			graphql.RunQuery(file5)
			graphql.RunQuery(file6)

		}
	}

	utils.Iterate(fn, utils.IterateOptions{
		Iterations: 100,
		Pace:       100,
	})
}
