package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

type Classification struct {
	ClassificationID uuid.UUID
	Name             string
	UpdatedAt        time.Time
	CreatedAt        time.Time
}

func SelectClassificationByID(classificationID uuid.UUID) (*Classification, error) {
	row := db.QueryRow(queries.SelectClassificationByIDQuery, classificationID)

	var classification Classification
	var updatedAt sql.NullInt64
	var createdAt int64

	cols := []interface{}{
		&classification.ClassificationID,
		&classification.Name,
		&updatedAt,
		&createdAt,
	}

	err := row.Scan(cols...)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		classification.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	classification.CreatedAt = time.UnixMilli(createdAt)

	return &classification, nil
}
