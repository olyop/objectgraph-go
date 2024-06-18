package database

import (
	"context"
	"database/sql"
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
	scanner := contactsRowScanner(rows)

	for rows.Next() {
		contact, err := scanner()
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func contactsRowScanner(scanner Scanner) func() (*Contact, error) {
	return func() (*Contact, error) {
		var contact Contact

		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []interface{}{
			&contact.ContactID,
			&contact.Value,
			&contact.Type,
			&updatedAt,
			&createdAt,
		}

		err := scanner.Scan(cols...)
		if err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			contact.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		contact.CreatedAt = time.UnixMilli(createdAt)

		return &contact, nil
	}
}

type Contact struct {
	ContactID uuid.UUID
	Value     string
	Type      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
