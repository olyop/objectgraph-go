package database

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"
)

//go:embed queries/select-brand-by-id.sql
var selectBrandByID string

type Brand struct {
	BrandID   uuid.UUID
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func SelectBrandByID(ctx context.Context, brandID uuid.UUID) (*Brand, error) {
	rows, err := db.QueryContext(ctx, selectBrandByID, brandID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		return brandRowMapper(rows), nil
	}

	return nil, nil
}
