package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

type Brand struct {
	BrandID   uuid.UUID
	Name      string
	CreatedAt time.Time
}

func SelectBrandByID(brandID uuid.UUID) (*Brand, error) {
	row := db.QueryRow(queries.SelectBrandByIDQuery, brandID)

	var brand Brand
	var createdAt int64

	cols := []interface{}{
		&brand.BrandID,
		&brand.Name,
		&createdAt,
	}

	err := row.Scan(cols...)
	if err != nil {
		return nil, err
	}

	brand.CreatedAt = time.UnixMilli(createdAt)

	return &brand, nil
}
