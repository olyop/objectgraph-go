package database

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	BrandID   uuid.UUID
	Name      string
	UpdatedAt *time.Time
	CreatedAt time.Time
}

type Category struct {
	CategoryID       uuid.UUID
	Name             string
	ClassificationID uuid.UUID
	UpdatedAt        *time.Time
	CreatedAt        time.Time
}

type Classification struct {
	ClassificationID uuid.UUID
	Name             string
	UpdatedAt        *time.Time
	CreatedAt        time.Time
}

type Contact struct {
	ContactID uuid.UUID
	Value     string
	Type      string
	UpdatedAt *time.Time
	CreatedAt time.Time
}

type Product struct {
	ProductID                 uuid.UUID
	Name                      string
	BrandID                   uuid.UUID
	Price                     *int
	ABV                       *int
	Volume                    *int
	PromotionDiscount         *int
	PromotionDiscountMultiple *int
	UpdatedAt                 *time.Time
	CreatedAt                 time.Time
}

type User struct {
	UserID    uuid.UUID
	UserName  string
	FirstName string
	LastName  string
	DOB       time.Time
	UpdatedAt *time.Time
	CreatedAt time.Time
}
