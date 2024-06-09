package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers/scalars"
)

func (*Resolver) GetProducts(ctx context.Context) ([]*ProductResolver, error) {
	return engine.ResolverList[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-products",
	})
}

func (*Resolver) GetProductByID(ctx context.Context, args struct{ ProductID scalars.UUID }) (*ProductResolver, error) {
	return engine.Resolver[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-product",
		RetrieverArgs: engine.RetrieverArgs{"productID": args.ProductID.UUID.String()},
	})
}
