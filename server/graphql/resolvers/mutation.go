package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
)

func (*Resolver) UpdateProductByID(ctx context.Context, args *UpdateProductByIDArgs) (*ProductResolver, error) {
	return graphqlops.Resolver[ProductResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveProductByID",
		RetrieverArgs: graphqlops.RetrieverArgs{"productID": args.Input.ProductID.Value},
	})
}

type UpdateProductByIDArgs struct {
	Input *UpdateProductByIDInput
}

type UpdateProductByIDInput struct {
	ProductID scalars.UUID
	Name      string
}

func (*Resolver) ClearCache() (bool, error) {
	err := graphqlops.ClearCache()
	if err != nil {
		return false, err
	}

	return true, nil
}
