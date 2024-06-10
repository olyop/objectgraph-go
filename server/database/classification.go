package database

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"
)

//go:embed queries/select-classification-by-id.sql
var selectClassificationByID string

type Classification struct {
	ClassificationID uuid.UUID
	Name             string
	UpdatedAt        time.Time
	CreatedAt        time.Time
}

func SelectClassificationByID(ctx context.Context, classificationID uuid.UUID) (*Classification, error) {
	rows, err := db.QueryContext(ctx, selectClassificationByID, classificationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		return classificationRowMapper(rows), nil
	}

	return nil, nil
}
