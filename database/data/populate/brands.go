package populate

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/olyop/objectgraph-go/data/database"
	"github.com/olyop/objectgraph-go/data/files"
)

func populateBrands(data *files.Data) (map[string]string, map[string][]string) {
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
			row = fmt.Sprintf("%s, ", values)
		} else {
			row = values
		}
		sql.WriteString(row)
	}

	os.WriteFile("/home/op/Downloads/brands.sql", []byte(sql.String()), 0644)

	sql.WriteString(" RETURNING brand_id, brand_name")

	rows, err := database.DB.Query(sql.String(), convertSetToArr(brandsSet)...)
	if err != nil {
		log.Default().Panic(err)
	}
	log.Default().Println("Populated brands")

	defer rows.Close()

	brands := brandsRowsMapper(rows)

	return brands, classificationsToBrands
}

func brandsRowsMapper(rows *sql.Rows) map[string]string {
	brands := make(map[string]string)

	for rows.Next() {
		var brandID string
		var brandName string

		err := rows.Scan(&brandID, &brandName)
		if err != nil {
			log.Default().Panic(err)
		}

		brands[brandName] = brandID
	}

	return brands
}
