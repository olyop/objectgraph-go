package resolvers

import (
	"context"

	"github.com/olyop/graphqlops-go/database"
	"github.com/olyop/graphqlops-go/graphql/scalars"
	"github.com/olyop/graphqlops-go/graphqlops"
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
