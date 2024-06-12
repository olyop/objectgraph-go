package files

import "embed"

//go:embed ContactTypes.txt
var contactTypesFile string

//go:embed Users.txt
var usersFile string

//go:embed Classifications.txt
var classificationFile string

//go:embed Brands
var brandsFolder embed.FS

//go:embed Categories
var categoriesFolder embed.FS

func Read() *Data {
	contactTypes := processFile(contactTypesFile)
	users := processFileTabSeperated(usersFile)
	classifications := processFile(classificationFile)
	classificationsToBrands := processFolderMap(brandsFolder, "Brands", classifications)
	classificationsToCategories := processFolderMap(categoriesFolder, "Categories", classifications)

	data := &Data{
		ContactTypes:               contactTypes,
		Users:                      users,
		Classifications:            classifications,
		ClassificationToBrands:     classificationsToBrands,
		ClassificationToCategories: classificationsToCategories,
	}

	return data
}

type Data struct {
	ContactTypes               []string
	Users                      [][]string
	Classifications            []string
	ClassificationToCategories map[string][]string // classification -> categories
	ClassificationToBrands     map[string][]string // classification -> brands
}
