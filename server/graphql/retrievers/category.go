package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
)

func (*Retrievers) RetrieveProductCategories(ctx context.Context, args graphqlops.RetrieverArgs) (any, error) {
	productID := args["productID"].(uuid.UUID)

	categories, err := database.SelectCategoriesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	r := make([]*resolvers.CategoryResolver, len(categories))
	for i := range categories {
		r[i] = mapToCategoryResolver(categories[i])
	}

	return r, nil
}

func mapToCategoryResolver(category *database.Category) *resolvers.CategoryResolver {
	return &resolvers.CategoryResolver{
		Category: category,

		CategoryID: scalars.NewUUID(category.CategoryID),
		Name:       category.Name,
		UpdatedAt:  scalars.NewNilTimestamp(category.UpdatedAt),
		CreatedAt:  scalars.NewTimestamp(category.CreatedAt),
	}
}
