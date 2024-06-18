package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
)

type CategoryResolver struct {
	Category *database.Category

	CategoryID scalars.UUID
	Name       string
	UpdatedAt  *scalars.Timestamp
	CreatedAt  scalars.Timestamp
}

func (r *CategoryResolver) Classification(ctx context.Context) (*ClassificationResolver, error) {
	return graphqlops.Resolver[ClassificationResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveClassificationByID",
		RetrieverArgs: graphqlops.RetrieverArgs{"classificationID": r.Category.ClassificationID},
	})
}
