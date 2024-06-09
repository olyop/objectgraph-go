package populate

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/import"
)

func populateBrands(data *importdata.Data) (map[string][]string, map[string]string) {
	classificationsToBrands := make(map[string][]string)

	brandsSet := make(map[string]struct{})

	for classification, classificationBrands := range data.ClassificationToBrands {
		for _, brand := range classificationBrands {
			if _, ok := brandsSet[brand]; ok {
				continue
			}

			brandsSet[brand] = struct{}{}

			classificationsToBrands[classification] = append(classificationsToBrands[classification], brand)
		}
	}

	var sql strings.Builder
	sql.WriteString("INSERT INTO brands (brand_name) VALUES ")

	for i := 1; i <= len(brandsSet); i++ {
		values := fmt.Sprintf("($%d)", i)

		var row string
		if i < len(brandsSet) {
			row = fmt.Sprintf("%s,", row)
		} else {
			row = values
		}
		sql.WriteString(row)
	}

	sql.WriteString(" RETURNING brand_id, brand_name")

	rows, err := database.DB.Query(sql.String(), convertSetToArr(brandsSet)...)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	brands := brandsRowsMapper(rows)

	log.Printf("Populated %d brands", len(brands))

	return classificationsToBrands, brands
}

func brandsRowsMapper(rows *sql.Rows) map[string]string {
	brands := make(map[string]string)

	for rows.Next() {
		var brandID string
		var brandName string

		err := rows.Scan(&brandID, &brandName)
		if err != nil {
			panic(err)
		}

		brands[brandName] = brandID
	}

	return brands
}

func clearBrands() {
	_, err := database.DB.Exec("DELETE FROM brands")
	if err != nil {
		panic(err)
	}
}
