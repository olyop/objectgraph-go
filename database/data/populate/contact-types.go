package populate

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/olyop/objectgraph-go/data/database"
	"github.com/olyop/objectgraph-go/data/files"
)

func populateContactTypes(data *files.Data) map[string]string {
	sql, params := createContactTypesQuery(data)

	rows, err := database.DB.Query(sql, params...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated contact_types")

	defer rows.Close()

	contactTypes := contactTypesRowsMapper(rows)

	return contactTypes
}

func createContactTypesQuery(data *files.Data) (string, []any) {
	var sql strings.Builder
	params := make([]string, len(data.ContactTypes))

	sql.WriteString("INSERT INTO contact_types (contact_type_name) VALUES ")

	for i := 0; i < len(data.ContactTypes); i++ {
		values := fmt.Sprintf("($%d)", i+1)

		var row string
		if i < len(data.ContactTypes)-1 {
			row = fmt.Sprintf("%s, ", values)
		} else {
			row = values
		}
		sql.WriteString(row)

		params[i] = data.ContactTypes[i]
	}

	sql.WriteString(" RETURNING contact_type_id, contact_type_name")

	return sql.String(), convertToInterfaceSlice(params)
}

func contactTypesRowsMapper(rows *sql.Rows) map[string]string {
	contactTypes := make(map[string]string, 0)

	for rows.Next() {
		var contactTypeID string
		var contactTypeName string

		err := rows.Scan(&contactTypeID, &contactTypeName)
		if err != nil {
			log.Default().Fatal(err)
		}

		contactTypes[contactTypeName] = contactTypeID
	}

	return contactTypes
}
