package resolvers

import "database/sql"

func handleSqlNullInt32(n sql.NullInt32) *int {
	if !n.Valid {
		return nil
	}

	value := int(n.Int32)

	return &value
}
