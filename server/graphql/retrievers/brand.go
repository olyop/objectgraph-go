package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
)

func (*Retrievers) RetrieveBrandByID(ctx context.Context, args graphqlops.RetrieverArgs) (any, error) {
	brandID := args["brandID"].(uuid.UUID)

	brand, err := database.SelectBrandByID(ctx, brandID)
	if err != nil {
		return nil, err
	}

	return mapToBrandResolver(brand), nil
}

func mapToBrandResolver(brand *database.Brand) *resolvers.BrandResolver {
	return &resolvers.BrandResolver{
		BrandID:   scalars.NewUUID(brand.BrandID),
		Name:      brand.Name,
		UpdatedAt: scalars.NewNilTimestamp(brand.UpdatedAt),
		CreatedAt: scalars.NewTimestamp(brand.CreatedAt),
	}
}
