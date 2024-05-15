package resolvers

import (
	"github.com/olyop/graphql-go/database"
	"github.com/olyop/graphql-go/resolvers/scalars"
)

type BrandResolver struct {
	brand *database.Brand
}

func (r *BrandResolver) BrandID() scalars.UUID {
	return scalars.UUID{UUID: r.brand.BrandID}
}

func (r *BrandResolver) Name() string {
	return r.brand.Name
}

func (r *BrandResolver) CreatedAt() scalars.Timestamp {
	return scalars.Timestamp{Time: r.brand.CreatedAt}
}
