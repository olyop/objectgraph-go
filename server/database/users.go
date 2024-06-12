package database

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database/queries"
)

func SelectTop1000Users(ctx context.Context) ([]*User, error) {
	rows, err := db.QueryContext(ctx, queries.SelectTop1000Users)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		users = append(users, userRowMapper(rows))
	}

	return users, nil
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
