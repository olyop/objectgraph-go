package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
)

type UserResolver struct {
	User *database.User

	UserID    scalars.UUID
	UserName  string
	FirstName string
	LastName  string
	DOB       scalars.Timestamp
	UpdatedAt *scalars.Timestamp
	CreatedAt scalars.Timestamp
}

func (r *UserResolver) Contacts(ctx context.Context) ([]*ContactResolver, error) {
	return graphqlops.ResolverList[ContactResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveUserContacts",
		RetrieverArgs: graphqlops.RetrieverArgs{"userID": r.User.UserID},
	})
}
