package resolvers

import (
	"time"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/resolvers/engine"
)

func (*Resolver) GetProducts() ([]*ProductResolver, error) {
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
