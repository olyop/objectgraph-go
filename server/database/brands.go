package database

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

func SelectBrandByID(ctx context.Context, brandID uuid.UUID) (*Brand, error) {
	rows, err := db.QueryContext(ctx, queries.SelectBrandByID, brandID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		return brandRowMapper(rows), nil
	}

	return nil, nil
}

type Brand struct {
	BrandID   uuid.UUID
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
