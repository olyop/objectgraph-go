package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

type Brand struct {
	BrandID   uuid.UUID
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func SelectBrandByID(brandID uuid.UUID) (*Brand, error) {
	row := db.QueryRow(queries.SelectBrandByIDQuery, brandID)

	var brand Brand
	var updatedAt sql.NullInt64
	var createdAt int64

	cols := []interface{}{
		&brand.BrandID,
		&brand.Name,
		&updatedAt,
		&createdAt,
	}

	err := row.Scan(cols...)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		brand.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	brand.CreatedAt = time.UnixMilli(createdAt)

	return &brand, nil
}
