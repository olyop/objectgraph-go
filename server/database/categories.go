package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database/queries"
)

func SelectCategoryByID(categoryID uuid.UUID) (*Category, error) {
	rows, err := db.Query(queries.SelectCategoryByID, categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	scanner := categoryRowScanner(rows)

	if rows.Next() {
		return scanner()
	}

	return nil, nil

}

func SelectCategoriesByIDs(categoryIDs []uuid.UUID) ([]*Category, error) {
	rows, err := db.Query(queries.SelectCategoriesByIDs, categoryIDs)
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

func SelectCategoriesByProductID(productID uuid.UUID) ([]*Category, error) {
	rows, err := db.Query(queries.SelectCategoriesByProductID, productID)
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

		cols := []any{
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
			value := time.UnixMilli(updatedAt.Int64)
			category.UpdatedAt = &value
		}
		category.CreatedAt = time.UnixMilli(createdAt)

		return &category, nil
	}
}
