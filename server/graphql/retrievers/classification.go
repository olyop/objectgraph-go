package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
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
