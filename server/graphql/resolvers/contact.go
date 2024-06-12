package resolvers

import "github.com/olyop/graphql-go/server/graphql/scalars"

type ContactResolver struct {
	ContactID scalars.UUID
	Value     string
	Type      string
	UpdatedAt *scalars.Timestamp
	CreatedAt scalars.Timestamp
}
