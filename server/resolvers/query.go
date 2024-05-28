package resolvers

import (
	"context"
	"time"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
)

func (*Resolver) GetProducts(ctx context.Context) ([]*ProductResolver, error) {
	return engine.Resolver(engine.ResolverOptions[[]*ProductResolver]{
		GroupKey: "query",
		Duration: time.Second * 15,
		CacheKey: "products",
		Retrieve: retrieveProducts,
	})
}

func retrieveProducts() ([]*ProductResolver, error) {
	products, err := database.SelectProducts()
	if err != nil {
		return nil, err
	}

	productsResolvers := make([]*ProductResolver, len(products))
	for i, product := range products {
		productsResolvers[i] = &ProductResolver{product}
	}

	return productsResolvers, nil
}
