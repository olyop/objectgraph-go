package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers/scalars"
)

func (*Resolver) UpdateProductByID(ctx context.Context, args *UpdateProductByIDArgs) (*ProductResolver, error) {
	return engine.Resolver[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-product",
		RetrieverArgs: engine.RetrieverArgs{"productID": args.Input.ProductID.UUID.String()},
	})
}

type UpdateProductByIDArgs struct {
	Input *UpdateProductByIDInput
}

type UpdateProductByIDInput struct {
	ProductID scalars.UUID
	Name      string
}
