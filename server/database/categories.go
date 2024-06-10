package database

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"
)

//go:embed queries/select-categories-by-product-id.sql
var selectCategoriesByProductID string

type Category struct {
	CategoryID       uuid.UUID
	Name             string
	ClassificationID uuid.UUID
	UpdatedAt        time.Time
	CreatedAt        time.Time
}

func SelectCategoriesByProductID(ctx context.Context, productID uuid.UUID) ([]*Category, error) {
	rows, err := db.QueryContext(ctx, selectCategoriesByProductID, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := make([]*Category, 0)

	for rows.Next() {
		categories = append(categories, categoryRowMapper(rows))
	}

	return categories, nil
}
