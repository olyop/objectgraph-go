package resolvers

import (
	"github.com/olyop/graphql-go/server/graphql/enums"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

type ContactResolver struct {
	ContactID   scalars.UUID
	Value       string
	ContactType enums.ContactType
	UpdatedAt   *scalars.Timestamp
	CreatedAt   scalars.Timestamp
}
