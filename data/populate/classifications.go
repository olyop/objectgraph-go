package populate

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/import"
)

func populateClassifications(data *importdata.Data) map[string]string {
	sql, params := createClassificationsQuery(data)

	rows, err := database.DB.Query(sql, params...)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	classifications := classificationsMap(rows)

	logClassifications(classifications)

	return classifications
}

func createClassificationsQuery(data *importdata.Data) (string, []interface{}) {
	var sql strings.Builder
	params := make([]string, len(data.Classifications))

	sql.WriteString("INSERT INTO classifications (classification_name) VALUES ")

	for i := 0; i < len(data.Classifications); i++ {
		row := fmt.Sprintf("($%d)", i+1)

		if i < len(data.Classifications)-1 {
			sql.WriteString(fmt.Sprintf("%s,", row))
		} else {
			sql.WriteString(row)
		}

		params[i] = data.Classifications[i]
	}

	sql.WriteString(" RETURNING classification_id, classification_name")

	return sql.String(), convertToInterfaceSlice(params)
}

func classificationsMap(rows *sql.Rows) map[string]string {
	classifications := make(map[string]string, 0)

	for rows.Next() {
		var classificationID string
		var classificationName string

		err := rows.Scan(&classificationID, &classificationName)
		if err != nil {
			panic(err)
		}

		classifications[classificationName] = classificationID
	}

	return classifications
}

func logClassifications(classifications map[string]string) {
	for classificationName, classificationID := range classifications {
		fmt.Printf("Added Classification: %v, %v\n", classificationID, classificationName)
	}
}

func clearClassifications() {
	_, err := database.DB.Exec("DELETE FROM classifications")
	if err != nil {
		panic(err)
	}
}
