package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphqlops-go/database/queries"
)

func SelectTop1000Users(ctx context.Context) ([]*User, error) {
	rows, err := db.QueryContext(ctx, queries.SelectTop1000Users)
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
			user.UpdatedAt = time.UnixMilli(updatedAt.Int64)
		}

		user.DOB = time.UnixMilli(dob)
		user.CreatedAt = time.UnixMilli(createdAt)

		return &user, nil
	}
}

type User struct {
	UserID    uuid.UUID
	UserName  string
	FirstName string
	LastName  string
	DOB       time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}
