package database

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

func SelectCategoriesByProductID(ctx context.Context, productID uuid.UUID) ([]*Category, error) {
	rows, err := db.QueryContext(ctx, queries.SelectCategoriesByProductID, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := make([]*Category, 0)
	scanner := categoryRowScanner(rows)

	for rows.Next() {
		category, err := scanner()
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func categoryRowScanner(scanner Scanner) func() (*Category, error) {
	return func() (*Category, error) {

		var category Category

		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []interface{}{
			&category.CategoryID,
			&category.Name,
			&category.ClassificationID,
			&updatedAt,
			&createdAt,
		}

		err := scanner.Scan(cols...)
		if err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			category.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		category.CreatedAt = time.UnixMilli(createdAt)

		return &category, nil
	}
}

type Category struct {
	CategoryID       uuid.UUID
	Name             string
	ClassificationID uuid.UUID
	UpdatedAt        time.Time
	CreatedAt        time.Time
}
