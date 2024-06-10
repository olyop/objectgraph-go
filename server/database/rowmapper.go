package database

import (
	"database/sql"
	"time"
)

func brandRowMapper(scanner Scanner) *Brand {
	var brand Brand

	var updatedAt sql.NullInt64
	var createdAt int64

	cols := []interface{}{
		&brand.BrandID,
		&brand.Name,
		&updatedAt,
		&createdAt,
	}

	err := scanner.Scan(cols...)
	if err != nil {
		return nil
	}

	if updatedAt.Valid {
		brand.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	brand.CreatedAt = time.UnixMilli(createdAt)

	return &brand
}

func categoryRowMapper(scanner Scanner) *Category {
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
		return nil
	}

	if updatedAt.Valid {
		category.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	category.CreatedAt = time.UnixMilli(createdAt)

	return &category
}

func classificationRowMapper(scanner Scanner) *Classification {
	var classification Classification

	var updatedAt sql.NullInt64
	var createdAt int64

	cols := []interface{}{
		&classification.ClassificationID,
		&classification.Name,
		&updatedAt,
		&createdAt,
	}

	err := scanner.Scan(cols...)
	if err != nil {
		return nil
	}

	if updatedAt.Valid {
		classification.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	classification.CreatedAt = time.UnixMilli(createdAt)

	return &classification
}

func productRowMapper(scanner Scanner) *Product {
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
		return nil
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

	return &product
}

type Scanner interface {
	Scan(dest ...interface{}) error
}
