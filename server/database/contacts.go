package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

func SelectContactsByUserID(ctx context.Context, userID uuid.UUID) ([]*Contact, error) {
	rows, err := db.QueryContext(ctx, queries.SelectContactsByUserID, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	contacts := make([]*Contact, 0)

	for rows.Next() {
		contacts = append(contacts, contactRowMapper(rows))
	}

	return contacts, nil
}

type Contact struct {
	ContactID uuid.UUID
	Value     string
	Type      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
