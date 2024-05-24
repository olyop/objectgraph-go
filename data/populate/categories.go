package populate

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/import"
)

func populateCategories(data *importdata.Data, classifications map[string]string) map[string]string {
	sql, params := createCategoriesQuery(data, classifications)

	rows, err := database.DB.Query(sql, params...)
	if err != nil {
		panic(err)
	}

	categories := constructCategories(rows)

	logCategories(categories)

	return categories
}

func createCategoriesQuery(data *importdata.Data, classifications map[string]string) (string, []interface{}) {
	var sql strings.Builder
	params := make([]string, 0)

	sql.WriteString("INSERT INTO categories (category_name, classification_id) VALUES ")

	i := 0
	for classification, categories := range data.ClassificationToCategories {
		for j, category := range categories {
			paramsLength := len(params) + 1

			row := fmt.Sprintf("($%d,$%d)", paramsLength, paramsLength+1)

			if j < len(categories)-1 {
				sql.WriteString(fmt.Sprintf("%s,", row))
			} else {
				sql.WriteString(row)
			}

			params = append(params, category, classifications[classification])
		}

		if i < len(data.ClassificationToCategories)-1 {
			sql.WriteString(",")
		}

		i++
	}

	sql.WriteString(" RETURNING category_id, category_name")

	return sql.String(), convertToInterfaceSlice(params)
}

func constructCategories(rows *sql.Rows) map[string]string {
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

func logCategories(categories map[string]string) {
	for categoryName, categoryID := range categories {
		fmt.Printf("Added Category: %v, %v\n", categoryID, categoryName)
	}
}

func clearCategories() {
	_, err := database.DB.Exec("DELETE FROM categories")
	if err != nil {
		panic(err)
	}
}
