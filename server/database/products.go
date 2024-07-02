package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database/queries"
)

func SelectProductByID(productID uuid.UUID) (*Product, error) {
	rows, err := db.Query(queries.SelectProductByID, productID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	scanner := productRowsScanner(rows)
	if rows.Next() {
		return scanner()
	}

	return nil, nil
}

func SelectProductsByIDs(productIDs []uuid.UUID) ([]*Product, error) {
	rows, err := db.Query(queries.SelectProductsByIDs, productIDs)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, len(productIDs))
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

func SelectTop1000Products() ([]*Product, error) {
	rows, err := db.Query(queries.SelectTop1000Products)
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

func productRowsScanner(scanner Scanner) func() (*Product, error) {
	return func() (*Product, error) {
		var product Product
		var price sql.NullInt32
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
			&price,
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

		if price.Valid {
			value := int(price.Int32)
			product.Price = &value
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
			value := time.UnixMilli(updatedAt.Int64)
			product.UpdatedAt = &value
		}
		product.CreatedAt = time.UnixMilli(createdAt)

		return &product, nil
	}
}
