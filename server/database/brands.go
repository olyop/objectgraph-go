package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/olyop/objectgraph/database/queries"
)

func SelectBrandByID(brandID uuid.UUID) (*Brand, error) {
	rows, err := db.Query(queries.SelectBrandByID, brandID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	scanner := brandRowMapper(rows)

	if rows.Next() {
		return scanner()
	}

	return nil, nil
}

func SelectBrandsByIDs(brandIDs []uuid.UUID) ([]*Brand, error) {
	rows, err := db.Query(queries.SelectBrandsByIDs, pq.Array(brandIDs))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	brands := make([]*Brand, len(brandIDs))
	scanner := brandRowMapper(rows)
	for rows.Next() {
		brand, err := scanner()
		if err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}

	return brands, nil
}

func brandRowMapper(scanner Scanner) func() (*Brand, error) {
	return func() (*Brand, error) {
		var brand Brand
		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []any{
			&brand.BrandID,
			&brand.Name,
			&updatedAt,
			&createdAt,
		}
		err := scanner.Scan(cols...)
		if err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			value := time.UnixMilli(updatedAt.Int64)
			brand.UpdatedAt = &value
		}
		brand.CreatedAt = time.UnixMilli(createdAt)

		return &brand, nil
	}
}
