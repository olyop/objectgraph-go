package database

import (
	"context"
	_ "embed"
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

func SelectClassificationByID(ctx context.Context, classificationID uuid.UUID) (*Classification, error) {
	rows, err := db.QueryContext(ctx, queries.SelectClassificationByID, classificationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		return classificationRowMapper(rows), nil
	}

	return nil, nil
}
