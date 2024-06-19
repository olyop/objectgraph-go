package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphqlops-go/database"
	"github.com/olyop/graphqlops-go/graphql/resolvers"
	"github.com/olyop/graphqlops-go/graphql/scalars"
	"github.com/olyop/graphqlops-go/graphqlops"
)

func (*Retrievers) RetrieveClassificationByID(ctx context.Context, args graphqlops.RetrieverArgs) (any, error) {
	classificationID := args["classificationID"].(uuid.UUID)

	classification, err := database.SelectClassificationByID(ctx, classificationID)
	if err != nil {
		return nil, err
	}

	return mapToClassificationResolver(classification), nil
}

func mapToClassificationResolver(classification *database.Classification) *resolvers.ClassificationResolver {
	return &resolvers.ClassificationResolver{
		ClassificationID: scalars.NewUUID(classification.ClassificationID),
		Name:             classification.Name,
		UpdatedAt:        scalars.NewNilTimestamp(classification.UpdatedAt),
		CreatedAt:        scalars.NewTimestamp(classification.CreatedAt),
	}
}
