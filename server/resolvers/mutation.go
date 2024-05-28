package resolvers

import (
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/resolvers/scalars"
)

func (*Resolver) UpdateProductByID(args *UpdateProductByIDArgs) (*ProductResolver, error) {
	productID := args.Input.ProductID.UUID

	return engine.Resolver(engine.ResolverOptions[*ProductResolver]{
		GroupKey: "product",
		CacheKey: productID.String(),
		Duration: time.Second * 15,
		Retrieve: retrieveProduct(productID),
	})
}

func retrieveProduct(productID uuid.UUID) func() (*ProductResolver, error) {
	return func() (*ProductResolver, error) {
		product, err := database.SelectProductByID(productID)
		if err != nil {
			return nil, err
		}

		return &ProductResolver{product}, nil
	}
}

type UpdateProductByIDArgs struct {
	Input *UpdateProductByIDInput
}

type UpdateProductByIDInput struct {
	ProductID scalars.UUID
	Name      string
}
