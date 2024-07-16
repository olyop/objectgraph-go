package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database/queries"
)

func SelectClassificationByID(classificationID uuid.UUID) (*Classification, error) {
	rows, err := db.Query(queries.SelectClassificationByID, classificationID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	scanner := classificationRowScanner(rows)
	if rows.Next() {
		return scanner()
	}

	return nil, nil
}

func SelectClassificationsByIDs(classificationIDs []uuid.UUID) ([]*Classification, error) {
	rows, err := db.Query(queries.SelectClassificationsByIDs, classificationIDs)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	classifications := make([]*Classification, 0)
	scanner := classificationRowScanner(rows)
	for rows.Next() {
		classification, err := scanner()
		if err != nil {
			return nil, err
		}

		classifications = append(classifications, classification)
	}

	return classifications, nil
}

func classificationRowScanner(scanner Scanner) func() (*Classification, error) {
	return func() (*Classification, error) {
		var classification Classification
		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []any{
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
			value := time.UnixMilli(updatedAt.Int64)
			classification.UpdatedAt = &value
		}
		classification.CreatedAt = time.UnixMilli(createdAt)

		return &classification, nil
	}
}
