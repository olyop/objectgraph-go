package database

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	BrandID   uuid.UUID
	Name      string
	CreatedAt time.Time
}

var selectBrandByIDQuery = `
	SELECT
		brands.brand_id,
		brands.name,
		brands.created_at
	FROM
		brands
	WHERE
		brands.brand_id = $1;
`

func SelectBrandByID(brandID uuid.UUID) (brand Brand, err error) {
	rows, err := db.Query(selectBrandByIDQuery, brandID)
	if err != nil {
		return brand, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(
			&brand.BrandID,
			&brand.Name,
			&brand.CreatedAt,
		)

		if err != nil {
			return brand, err
		}
	}

	return brand, nil
}
