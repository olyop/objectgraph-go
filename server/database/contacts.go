package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database/queries"
)

func SelectContactsByUserID(userID uuid.UUID) ([]*Contact, error) {
	rows, err := db.Query(queries.SelectContactsByUserID, userID)
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
			value := time.UnixMilli(updatedAt.Int64)
			contact.UpdatedAt = &value
		}
		contact.CreatedAt = time.UnixMilli(createdAt)

		return &contact, nil
	}
}
