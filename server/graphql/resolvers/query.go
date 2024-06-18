package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
)

func (*Resolver) GetUsers(ctx context.Context) ([]*UserResolver, error) {
	return graphqlops.ResolverList[UserResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveTop1000Users",
	})
}

func (*Resolver) GetProducts(ctx context.Context) ([]*ProductResolver, error) {
	return graphqlops.ResolverList[ProductResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveTop1000Products",
	})
}

func (*Resolver) GetProductByID(ctx context.Context, args GetProductByIDArgs) (*ProductResolver, error) {
	return graphqlops.Resolver[ProductResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveProductByID",
		RetrieverArgs: graphqlops.RetrieverArgs{"productID": args.ProductID.Value},
	})
}

type GetProductByIDArgs struct {
	ProductID scalars.UUID
}
