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
	populateProducts(data, classificationsToBrands, brands, categories)

	fmt.Println(len(classifications))
	fmt.Println(len(categories))
	fmt.Println(len(classificationsToBrands))
	fmt.Println(len(brands))
}
