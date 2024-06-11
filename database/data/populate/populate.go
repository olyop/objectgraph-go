package populate

import (
	"log"

	"github.com/olyop/graphql-go/data/files"
)

func Execute(data *files.Data) {
	clearUsers()
	clearCategories()
	clearClassifications()
	clearProducts()
	clearBrands()

	users := populateUsers(data)
	classifications := populateClassifications(data)
	categories := populateCategories(data, classifications)
	brands, classificationsToBrands := populateBrands(data)
	products := populateProducts(data, brands, categories, classificationsToBrands)

	log.Default().Printf("Populated %d users\n", len(users))
	log.Default().Printf("Populated %d classifications\n", len(classifications))
	log.Default().Printf("Populated %d categories\n", len(categories))
	log.Default().Printf("Populated %d brands\n", len(brands))
	log.Default().Printf("Populated %d products\n", len(products))
}
