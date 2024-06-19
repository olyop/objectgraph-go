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

func SelectClassificationByID(ctx context.Context, classificationID uuid.UUID) (*Classification, error) {
	rows, err := db.QueryContext(ctx, queries.SelectClassificationByID, classificationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	scanner := classificationRowScanner(rows)

	if rows.Next() {
		return scanner()
	}

	return nil, fmt.Errorf("classification with id %s not found", classificationID)
}

func classificationRowScanner(scanner Scanner) func() (*Classification, error) {
	return func() (*Classification, error) {

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
			return nil, err
		}

		if updatedAt.Valid {
			classification.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		classification.CreatedAt = time.UnixMilli(createdAt)

		return &classification, nil
	}
}

type Classification struct {
	ClassificationID uuid.UUID
	Name             string
	UpdatedAt        time.Time
	CreatedAt        time.Time
}
