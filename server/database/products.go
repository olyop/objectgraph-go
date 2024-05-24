package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

type Product struct {
	ProductID uuid.UUID
	Name      string
	BrandID   uuid.UUID
	UpdatedAt time.Time
	CreatedAt time.Time
	Price     int64
	ABV       sql.NullFloat64
	Volume    sql.NullInt64
}

func SelectProducts() ([]*Product, error) {
	rows, err := db.Query(queries.SelectProductsQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		var product Product
		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []interface{}{
			&product.ProductID,
			&product.Name,
			&product.BrandID,
			&product.Price,
			&product.ABV,
			&product.Volume,
			&updatedAt,
			&createdAt,
		}

		err := rows.Scan(cols...)
		if err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			product.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		product.CreatedAt = time.UnixMilli(createdAt)

		products = append(products, &product)
	}

	return products, nil
}

func SelectProductByID(productID uuid.UUID) (*Product, error) {
	row := db.QueryRow(queries.SelectProductByID, productID)

	var product Product
	var updatedAt sql.NullInt64
	var createdAt int64

	cols := []interface{}{
		&product.ProductID,
		&product.Name,
		&product.BrandID,
		&product.Price,
		&product.ABV,
		&product.Volume,
		&updatedAt,
		&createdAt,
	}

	err := row.Scan(cols...)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		product.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	product.CreatedAt = time.UnixMilli(createdAt)

	return &product, nil
}
