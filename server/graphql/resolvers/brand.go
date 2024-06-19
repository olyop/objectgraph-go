package resolvers

import (
	"github.com/olyop/graphqlops-go/graphql/scalars"
)

type BrandResolver struct {
	BrandID   scalars.UUID
	Name      string
	UpdatedAt *scalars.Timestamp
	CreatedAt scalars.Timestamp
}
