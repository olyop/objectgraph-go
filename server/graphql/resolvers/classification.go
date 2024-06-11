package resolvers

import (
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

type ClassificationResolver struct {
	ClassificationID scalars.UUID
	Name             string
	UpdatedAt        *scalars.Timestamp
	CreatedAt        scalars.Timestamp
}
