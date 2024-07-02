package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database/queries"
)

func SelectUserByID(userID uuid.UUID) (*User, error) {
	rows, err := db.Query(queries.SelectUserByID, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	scanner := userRowScanner(rows)
	if rows.Next() {
		return scanner()
	}

	return nil, nil
}

func SelectUsersByIDs(userIDs []uuid.UUID) ([]*User, error) {
	rows, err := db.Query(queries.SelectUsersByIDs, userIDs)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User, len(userIDs))
	scanner := userRowScanner(rows)
	for rows.Next() {
		user, err := scanner()
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func SelectTop1000Users() ([]*User, error) {
	rows, err := db.Query(queries.SelectTop1000Users)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User, 0)
	scanner := userRowScanner(rows)
	for rows.Next() {
		user, err := scanner()
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func userRowScanner(scanner Scanner) func() (*User, error) {
	return func() (*User, error) {
		var user User
		var dob int64
		var updatedAt sql.NullInt64
		var createdAt int64

		cols := []interface{}{
			&user.UserID,
			&user.UserName,
			&user.FirstName,
			&user.LastName,
			&dob,
			&updatedAt,
			&createdAt,
		}
		err := scanner.Scan(cols...)
		if err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			value := time.UnixMilli(updatedAt.Int64)
			user.UpdatedAt = &value
		}
		user.DOB = time.UnixMilli(dob)
		user.CreatedAt = time.UnixMilli(createdAt)

		return &user, nil
	}
}
