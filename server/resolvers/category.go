package resolvers

import (
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/resolvers/scalars"
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

func (r *CategoryResolver) Classification() (*ClassificationResolver, error) {
	return engine.Resolver(engine.ResolverOptions[*ClassificationResolver]{
		GroupKey: "category-classifications",
		Duration: 15 * time.Second,
		CacheKey: r.Category.ClassificationID.String(),
		Retrieve: classificationRetriever(r.Category.ClassificationID),
	})
}

func (r *CategoryResolver) UpdatedAt() *scalars.Timestamp {
	return scalars.NewNillTimestamp(r.Category.UpdatedAt)
}

func (r *CategoryResolver) CreatedAt() scalars.Timestamp {
	return scalars.NewTimestamp(r.Category.CreatedAt)
}

func classificationRetriever(classificationID uuid.UUID) func() (*ClassificationResolver, error) {
	return func() (*ClassificationResolver, error) {
		classification, err := database.SelectClassificationByID(classificationID)
		if err != nil {
			return nil, err
		}

		return &ClassificationResolver{classification}, nil
	}
}
