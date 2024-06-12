package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func (*Resolver) GetUsers(ctx context.Context) ([]*UserResolver, error) {
	return engine.ResolverList[UserResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "retrieve-top-1000-users",
	})
}

func (*Resolver) GetProducts(ctx context.Context) ([]*ProductResolver, error) {
	return engine.ResolverList[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "retrieve-top-1000-products",
	})
}

func (*Resolver) GetProductByID(ctx context.Context, args GetProductByIDArgs) (*ProductResolver, error) {
	return engine.Resolver[ProductResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "retrieve-product-by-id",
		RetrieverArgs: engine.RetrieverArgs{"productID": args.ProductID.Value.String()},
	})
}

type GetProductByIDArgs struct {
	ProductID scalars.UUID
}
