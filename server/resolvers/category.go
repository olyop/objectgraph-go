package resolvers

import (
	"github.com/olyop/graphql-go/database"
	"github.com/olyop/graphql-go/resolvers/scalars"
)

type CategoryResolver struct {
	category *database.Category
}

func (r *CategoryResolver) CategoryID() scalars.UUID {
	return scalars.UUID{UUID: r.category.CategoryID}
}

func (r *CategoryResolver) Name() string {
	return r.category.Name
}

func (r *CategoryResolver) CreatedAt() scalars.Timestamp {
	return scalars.Timestamp{Time: r.category.CreatedAt}
}
