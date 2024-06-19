package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphqlops-go/database"
	"github.com/olyop/graphqlops-go/graphql/resolvers"
	"github.com/olyop/graphqlops-go/graphql/scalars"
	"github.com/olyop/graphqlops-go/graphqlops"
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

	return &r, nil
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
