package resolvers

import (
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/resolvers/scalars"
)

type BrandResolver struct {
	*database.Brand
}

func (r *BrandResolver) BrandID() scalars.UUID {
	return scalars.NewUUID(r.Brand.BrandID)
}

func (r *BrandResolver) Name() string {
	return r.Brand.Name
}

func (r *BrandResolver) UpdatedAt() *scalars.Timestamp {
	return scalars.NewNillTimestamp(r.Brand.UpdatedAt)
}

func (r *BrandResolver) CreatedAt() scalars.Timestamp {
	return scalars.NewTimestamp(r.Brand.CreatedAt)
}
