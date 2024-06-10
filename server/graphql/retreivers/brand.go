package retreivers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func GetBrandByID(ctx context.Context, args engine.RetrieverArgs) (any, error) {
	brandID, err := uuid.Parse(args["brandID"])
	if err != nil {
		return nil, err
	}

	brand, err := database.SelectBrandByID(ctx, brandID)
	if err != nil {
		return nil, err
	}

	return mapToBrandResolver(brand), nil
}

func mapToBrandResolver(brand *database.Brand) *resolvers.BrandResolver {
	if brand == nil {
		return nil
	}

	return &resolvers.BrandResolver{
		BrandID:   scalars.NewUUID(brand.BrandID),
		Name:      brand.Name,
		UpdatedAt: scalars.NewNilTimestamp(brand.UpdatedAt),
		CreatedAt: scalars.NewTimestamp(brand.CreatedAt),
	}
}
