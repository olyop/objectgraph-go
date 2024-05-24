package populate

import (
	"database/sql"
	"fmt"
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
		row := fmt.Sprintf("($%d)", i)

		if i < len(brandsSet) {
			sql.WriteString(fmt.Sprintf("%s,", row))
		} else {
			sql.WriteString(row)
		}
	}

	sql.WriteString(" RETURNING brand_id, brand_name")

	rows, err := database.DB.Query(sql.String(), convertSetToArr(brandsSet)...)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	brands := constructBrands(rows)

	logBrands(brands)

	return classificationsToBrands, brands
}

func constructBrands(rows *sql.Rows) map[string]string {
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

func logBrands(brands map[string]string) {
	for brandName, brandID := range brands {
		fmt.Printf("Added Brand: %v, %v\n", brandID, brandName)
	}
}

func clearBrands() {
	_, err := database.DB.Exec("DELETE FROM brands")
	if err != nil {
		panic(err)
	}
}
