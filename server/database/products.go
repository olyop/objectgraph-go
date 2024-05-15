package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/utils"
)

type Product struct {
	ProductID uuid.UUID
	Name      string
	BrandID   uuid.UUID
	PriceID   uuid.UUID
	CreatedAt time.Time
	Price     int
	ABV       int
	Volume    int
}

var selectProductsQuery = utils.ReadFile("database", "select-products.sql")

func SelectProducts() ([]Product, error) {
	rows, err := db.Query(selectProductsQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]Product, 0)

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ProductID,
			&product.Name,
			&product.BrandID,
			&product.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
