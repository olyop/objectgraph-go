package database

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"
)

//go:embed queries/select-top-1000-products.sql
var selectTop1000Products string

//go:embed queries/select-product-by-id.sql
var selectProductByID string

type Product struct {
	ProductID                 uuid.UUID
	Name                      string
	BrandID                   uuid.UUID
	Price                     int
	ABV                       *int
	Volume                    *int
	PromotionDiscount         *int
	PromotionDiscountMultiple *int
	UpdatedAt                 time.Time
	CreatedAt                 time.Time
}

func SelectProducts(ctx context.Context) ([]*Product, error) {
	rows, err := db.QueryContext(ctx, selectTop1000Products)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		products = append(products, productRowMapper(rows))
	}

	return products, nil
}

func SelectProductByID(ctx context.Context, productID uuid.UUID) (*Product, error) {
	rows, err := db.QueryContext(ctx, selectProductByID, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		return productRowMapper(rows), nil
	}

	return nil, nil
}
