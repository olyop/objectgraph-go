package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func RetreiveProductCategories(ctx context.Context, args engine.RetrieverArgs) (any, error) {
	productID, err := uuid.Parse(args["productID"])
	if err != nil {
		return nil, err
	}

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
	if category == nil {
		return nil
	}

	return &resolvers.CategoryResolver{
		Category: category,

		CategoryID: scalars.NewUUID(category.CategoryID),
		Name:       category.Name,
		UpdatedAt:  scalars.NewNilTimestamp(category.UpdatedAt),
		CreatedAt:  scalars.NewTimestamp(category.CreatedAt),
	}
}
