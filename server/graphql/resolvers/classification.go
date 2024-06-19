package resolvers

import (
	"github.com/olyop/graphqlops-go/graphql/scalars"
)

type ClassificationResolver struct {
	ClassificationID scalars.UUID
	Name             string
	UpdatedAt        *scalars.Timestamp
	CreatedAt        scalars.Timestamp
}
