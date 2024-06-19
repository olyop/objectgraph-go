package resolvers

import (
	"github.com/olyop/graphqlops-go/graphql/enums"
	"github.com/olyop/graphqlops-go/graphql/scalars"
)

type ContactResolver struct {
	ContactID   scalars.UUID
	Value       string
	ContactType enums.ContactType
	UpdatedAt   *scalars.Timestamp
	CreatedAt   scalars.Timestamp
}
