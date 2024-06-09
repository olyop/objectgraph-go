package resolvers

import (
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/resolvers/scalars"
)

type ClassificationResolver struct {
	*database.Classification
}

func (r *ClassificationResolver) ClassificationID() scalars.UUID {
	return scalars.NewUUID(r.Classification.ClassificationID)
}

func (r *ClassificationResolver) Name() string {
	return r.Classification.Name
}

func (r *ClassificationResolver) UpdatedAt() *scalars.Timestamp {
	return scalars.NewNillTimestamp(r.Classification.UpdatedAt)
}

func (r *ClassificationResolver) CreatedAt() scalars.Timestamp {
	return scalars.NewTimestamp(r.Classification.CreatedAt)
}
