package database

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	CategoryID uuid.UUID
	Name       string
	CreatedAt  time.Time
}

const selectCategoriesByProductIDQuery = `
	SELECT
		categories.category_id,
`
