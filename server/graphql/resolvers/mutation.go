package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func (*Resolver) UpdateProductByID(ctx context.Context, args *UpdateProductByIDArgs) (*ProductResolver, error) {
	return engine.Resolver[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "retrieve-product-by-id",
		RetrieverArgs: engine.RetrieverArgs{"productID": args.Input.ProductID.Value.String()},
	})
}

type UpdateProductByIDArgs struct {
	Input *UpdateProductByIDInput
}

type UpdateProductByIDInput struct {
	ProductID scalars.UUID
	Name      string
}
