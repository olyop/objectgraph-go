package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

type CategoryResolver struct {
	Category *database.Category

	CategoryID scalars.UUID
	Name       string
	UpdatedAt  *scalars.Timestamp
	CreatedAt  scalars.Timestamp
}

func (r *CategoryResolver) Classification(ctx context.Context) (*ClassificationResolver, error) {
	return engine.Resolver[ClassificationResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-classification-by-id",
		RetrieverArgs: engine.RetrieverArgs{"classificationID": r.Category.ClassificationID.String()},
	})
}
