package populate

import (
	"fmt"

	"github.com/olyop/graphql-go/data/import"
)

func Execute(data *importdata.Data) {
	clearCategories()
	clearClassifications()
	clearProducts()
	clearBrands()

	classifications := populateClassifications(data)
	categories := populateCategories(data, classifications)
	classificationsToBrands, brands := populateBrands(data)
	products := populateProducts(data, classificationsToBrands, brands, categories)

	fmt.Println("Classifications:", len(classifications))
	fmt.Println("Categories:", len(categories))
	fmt.Println("Brands:", len(brands))
	fmt.Println("Products:", len(products))
}
