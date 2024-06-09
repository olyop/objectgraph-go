package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers/scalars"
)

type CategoryResolver struct {
	*database.Category
}

func (r *CategoryResolver) CategoryID() scalars.UUID {
	return scalars.NewUUID(r.Category.CategoryID)
}

func (r *CategoryResolver) Name() string {
	return r.Category.Name
}

func (r *CategoryResolver) UpdatedAt() *scalars.Timestamp {
	return scalars.NewNillTimestamp(r.Category.UpdatedAt)
}

func (r *CategoryResolver) CreatedAt() scalars.Timestamp {
	return scalars.NewTimestamp(r.Category.CreatedAt)
}

func (r *CategoryResolver) Classification(ctx context.Context) (*ClassificationResolver, error) {
	return engine.Resolver[ClassificationResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-classification",
		RetrieverArgs: engine.RetrieverArgs{"classificationID": r.Category.ClassificationID.String()},
	})
}
