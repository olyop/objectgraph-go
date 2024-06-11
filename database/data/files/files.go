package files

import "embed"

//go:embed Users.txt
var usersFile string

//go:embed Classifications.txt
var classificationFile string

//go:embed Brands
var brandsFolder embed.FS

//go:embed Categories
var categoriesFolder embed.FS

func Read() *Data {
	users := processFileTabSeperated(usersFile)
	classifications := processFile(classificationFile)
	classificationsToBrands := processFolderMap(brandsFolder, classifications)
	classificationsToCategories := processFolderMap(categoriesFolder, classifications)

	data := &Data{
		Users:                      users,
		Classifications:            classifications,
		ClassificationToBrands:     classificationsToBrands,
		ClassificationToCategories: classificationsToCategories,
	}

	return data
}

type Data struct {
	Users                      [][]string
	Classifications            []string
	ClassificationToCategories map[string][]string // classification -> categories
	ClassificationToBrands     map[string][]string // classification -> brands
}
