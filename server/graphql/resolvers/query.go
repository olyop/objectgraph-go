package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func (*Resolver) GetProducts(ctx context.Context) ([]*ProductResolver, error) {
	return engine.ResolverList[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-products",
	})
}

func (*Resolver) GetProductByID(ctx context.Context, args GetProductByIDArgs) (*ProductResolver, error) {
	return engine.Resolver[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-product-by-id",
		RetrieverArgs: engine.RetrieverArgs{"productID": args.ProductID.Value.String()},
	})
}

type GetProductByIDArgs struct {
	ProductID scalars.UUID
}
