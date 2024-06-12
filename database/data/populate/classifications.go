package populate

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/files"
)

func populateClassifications(data *files.Data) map[string]string {
	sql, params := createClassificationsQuery(data)

	rows, err := database.DB.Query(sql, params...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated classifications")

	defer rows.Close()

	classifications := classificationsRowsMapper(rows)

	return classifications
}

func createClassificationsQuery(data *files.Data) (string, []interface{}) {
	var sql strings.Builder
	params := make([]string, len(data.Classifications))

	sql.WriteString("INSERT INTO classifications (classification_name) VALUES ")

	for i := 0; i < len(data.Classifications); i++ {
		values := fmt.Sprintf("($%d)", i+1)

		var row string
		if i < len(data.Classifications)-1 {
			row = fmt.Sprintf("%s, ", values)
		} else {
			row = values
		}
		sql.WriteString(row)

		params[i] = data.Classifications[i]
	}

	sql.WriteString(" RETURNING classification_id, classification_name")

	return sql.String(), convertToInterfaceSlice(params)
}

func classificationsRowsMapper(rows *sql.Rows) map[string]string {
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
