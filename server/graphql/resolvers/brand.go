package resolvers

import (
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

type BrandResolver struct {
	BrandID   scalars.UUID
	Name      string
	UpdatedAt *scalars.Timestamp
	CreatedAt scalars.Timestamp
}
