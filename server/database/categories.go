package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

type Category struct {
	CategoryID uuid.UUID
	Name       string
	UpdatedAt  time.Time
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
		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []interface{}{
			&category.CategoryID,
			&category.Name,
			&updatedAt,
			&createdAt,
		}

		err := rows.Scan(cols...)
		if err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			category.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		category.CreatedAt = time.UnixMilli(createdAt)

		categories = append(categories, &category)
	}

	return categories, nil
}
