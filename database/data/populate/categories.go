package populate

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/olyop/objectgraph-go/data/database"
	"github.com/olyop/objectgraph-go/data/files"
)

func populateCategories(data *files.Data, classifications map[string]string) map[string]string {
	sql, params := createCategoriesQuery(data, classifications)

	rows, err := database.DB.Query(sql, params...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated categories")

	categories := categoriesRowsMapper(rows)

	return categories
}

func createCategoriesQuery(data *files.Data, classifications map[string]string) (string, []any) {
	var sql strings.Builder
	paramsIndex := 0
	params := initializeParams(data)

	sql.WriteString("INSERT INTO categories (category_name, classification_id) VALUES ")

	i := 0
	for classification, categories := range data.ClassificationToCategories {
		params[paramsIndex] = classifications[classification]
		classificationIndex := paramsIndex
		paramsIndex++

		for categoryIndex, category := range categories {
			values := fmt.Sprintf("($%d, $%d)", paramsIndex+1, classificationIndex+1)

			var row string
			if categoryIndex < len(categories)-1 {
				row = fmt.Sprintf("%s, ", values)
			} else {
				row = values
			}
			sql.WriteString(row)

			params[paramsIndex] = category
			paramsIndex++
		}

		if i < len(data.ClassificationToCategories)-1 {
			sql.WriteString(",")
		}

		i++
	}

	sql.WriteString(" RETURNING category_id, category_name")

	return sql.String(), convertToInterfaceSlice(params)
}

func initializeParams(data *files.Data) []string {
	count := 0

	count += len(data.ClassificationToCategories)

	for _, categories := range data.ClassificationToCategories {
		count += len(categories)
	}

	return make([]string, count)
}

func categoriesRowsMapper(rows *sql.Rows) map[string]string {
	categories := make(map[string]string)

	for rows.Next() {
		var categoryID string
		var categoryName string

		err := rows.Scan(&categoryID, &categoryName)
		if err != nil {
			panic(err)
		}

		categories[categoryName] = categoryID
	}

	return categories
}
