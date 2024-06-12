package retrievers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func RetreiveTop1000Users(ctx context.Context, args engine.RetrieverArgs) (any, error) {
	users, err := database.SelectTop1000Users(ctx)
	if err != nil {
		return nil, err
	}

	r := make([]*resolvers.UserResolver, len(users))
	for i, user := range users {
		r[i] = mapToUserResolver(user)
	}

	return &r, nil
}

func mapToUserResolver(user *database.User) *resolvers.UserResolver {
	if user == nil {
		return nil
	}

	return &resolvers.UserResolver{
		User: user,

		UserID:    scalars.NewUUID(user.UserID),
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		DOB:       scalars.NewTimestamp(user.DOB),
		UpdatedAt: scalars.NewNilTimestamp(user.UpdatedAt),
		CreatedAt: scalars.NewTimestamp(user.CreatedAt),
	}
}
