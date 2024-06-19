package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphqlops-go/database/queries"
)

func SelectTop1000Products(ctx context.Context) ([]*Product, error) {
	rows, err := db.QueryContext(ctx, queries.SelectTop1000Products)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, 0)
	scanner := productRowsScanner(rows)

	for rows.Next() {
		product, err := scanner()
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func SelectProductByID(ctx context.Context, productID uuid.UUID) (*Product, error) {
	rows, err := db.QueryContext(ctx, queries.SelectProductByID, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	scanner := productRowsScanner(rows)

	if rows.Next() {
		return scanner()
	}

	return nil, fmt.Errorf("product with id %s not found", productID)
}

func productRowsScanner(scanner Scanner) func() (*Product, error) {
	return func() (*Product, error) {
		var product Product

		var abv sql.NullInt32
		var volume sql.NullInt32
		var promotionDiscount sql.NullInt32
		var promotionDiscountMultiple sql.NullInt32
		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []interface{}{
			&product.ProductID,
			&product.Name,
			&product.BrandID,
			&product.Price,
			&abv,
			&volume,
			&promotionDiscount,
			&promotionDiscountMultiple,
			&updatedAt,
			&createdAt,
		}

		err := scanner.Scan(cols...)
		if err != nil {
			return nil, err
		}

		if abv.Valid {
			value := int(abv.Int32)
			product.ABV = &value
		}

		if volume.Valid {
			value := int(volume.Int32)
			product.Volume = &value
		}

		if promotionDiscount.Valid {
			value := int(promotionDiscount.Int32)
			product.PromotionDiscount = &value
		}

		if promotionDiscountMultiple.Valid {
			value := int(promotionDiscountMultiple.Int32)
			product.PromotionDiscountMultiple = &value
		}

		if updatedAt.Valid {
			product.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		product.CreatedAt = time.UnixMilli(createdAt)

		return &product, nil
	}
}

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
