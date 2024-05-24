package importdata

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/olyop/graphql-go/data/database"
)

type Data struct {
	Classifications            []string
	ClassificationToCategories map[string][]string // classification -> categories
	ClassificationToBrands     map[string][]string // classification -> brands
}

func Read() *Data {
	classifications := processFile("files/Classifications.txt")

	return &Data{
		Classifications:            classifications,
		ClassificationToBrands:     processFolderMap("files/Brands", classifications),
		ClassificationToCategories: processFolderMap("files/Categories", classifications),
	}
}

func processFolderMap(folderPath string, classificiations []string) map[string][]string {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		panic(err)
	}

	folderMap := make(map[string][]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := strings.TrimSuffix(filepath.Base(file.Name()), filepath.Base(filepath.Ext(file.Name())))

		for _, classification := range classificiations {
			if !strings.EqualFold(classification, fileName) {
				continue
			}

			folderMap[fileName] = processFile(filepath.Join(folderPath, file.Name()))
		}
	}

	return folderMap
}

func processFile(path string) []string {
	contents, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(contents), "\n")
}

func retrieveClassificationID(name string) string {
	var classificationID string

	row := database.DB.QueryRow("SELECT classification_id FROM classifications WHERE name = $1", name)

	err := row.Scan(&classificationID)
	if err != nil {
		panic(err)
	}

	return classificationID
}
