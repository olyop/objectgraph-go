package resolvers

import (
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/resolvers/scalars"
)

type CategoryResolver struct {
	*database.Category
}

func (r *CategoryResolver) CategoryID() scalars.UUID {
	return scalars.UUID{UUID: r.Category.CategoryID}
}

func (r *CategoryResolver) Name() string {
	return r.Category.Name
}

func (r *CategoryResolver) UpdatedAt() *scalars.Timestamp {
	if r.Category.UpdatedAt.IsZero() {
		return nil
	}

	return &scalars.Timestamp{Time: r.Category.UpdatedAt}
}

func (r *CategoryResolver) CreatedAt() scalars.Timestamp {
	return scalars.Timestamp{Time: r.Category.CreatedAt}
}
