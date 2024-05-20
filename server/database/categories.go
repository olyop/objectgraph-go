package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

type Category struct {
	CategoryID uuid.UUID
	Name       string
	CreatedAt  time.Time
}

func SelectCategoriesByProductID(productID uuid.UUID) ([]*Category, error) {
	rows, err := db.Query(queries.SelectCategoriesByProductID, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := make([]*Category, 0)

	for rows.Next() {
		var category Category
		var createdAt int64

		cols := []interface{}{
			&category.CategoryID,
			&category.Name,
			&createdAt,
		}

		err := rows.Scan(cols...)
		if err != nil {
			return nil, err
		}

		category.CreatedAt = time.UnixMilli(createdAt)

		categories = append(categories, &category)
	}

	return categories, nil
}
