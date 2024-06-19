package resolvers

import (
	"context"

	"github.com/olyop/graphqlops-go/database"
	"github.com/olyop/graphqlops-go/graphql/scalars"
	"github.com/olyop/graphqlops-go/graphqlops"
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
