package retreivers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func GetClassificationByID(ctx context.Context, args engine.RetrieverArgs) (any, error) {
	classificationID, err := uuid.Parse(args["classificationID"])
	if err != nil {
		return nil, err
	}

	classification, err := database.SelectClassificationByID(ctx, classificationID)
	if err != nil {
		return nil, err
	}

	return mapToClassificationResolver(classification), nil
}

func mapToClassificationResolver(classification *database.Classification) *resolvers.ClassificationResolver {
	if classification == nil {
		return nil
	}

	return &resolvers.ClassificationResolver{
		ClassificationID: scalars.NewUUID(classification.ClassificationID),
		Name:             classification.Name,
		UpdatedAt:        scalars.NewNilTimestamp(classification.UpdatedAt),
		CreatedAt:        scalars.NewTimestamp(classification.CreatedAt),
	}
}
