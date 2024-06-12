package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/scalars"
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
	return engine.ResolverList[ContactResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "retrieve-user-contacts",
		RetrieverArgs: engine.RetrieverArgs{"userID": r.User.UserID.String()},
	})
}
