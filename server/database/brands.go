package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphqlops-go/database/queries"
)

func SelectBrandByID(ctx context.Context, brandID uuid.UUID) (*Brand, error) {
	rows, err := db.QueryContext(ctx, queries.SelectBrandByID, brandID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	scanner := brandRowMapper(rows)

	if rows.Next() {
		return scanner()
	}

	return nil, fmt.Errorf("brand with id %s not found", brandID)
}

func brandRowMapper(scanner Scanner) func() (*Brand, error) {
	return func() (*Brand, error) {
		var brand Brand

		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []interface{}{
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
			brand.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		brand.CreatedAt = time.UnixMilli(createdAt)

		return &brand, nil
	}
}

type Brand struct {
	BrandID   uuid.UUID
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
