package resolvers

import "database/sql"

func handleSqlNullInt32(n sql.NullInt32) *int32 {
	if !n.Valid {
		return nil
	}

	value := int32(n.Int32)

	return &value
}

func handleSqlNullInt64(n sql.NullInt64) *int64 {
	if !n.Valid {
		return nil
	}

	value := int64(n.Int64)

	return &value
}

func handleSqlNullFloat64(n sql.NullFloat64) *float64 {
	if !n.Valid {
		return nil
	}

	value := float64(n.Float64)

	return &value
}
