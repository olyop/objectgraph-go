package resolvers

import (
	"fmt"

	"github.com/olyop/graphql-go/database"
)

func (*Resolver) GetProducts() ([]*ProductResolver, error) {
	products, err := database.SelectProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to select products: %w", err)
	}

	productsResolvers := make([]*ProductResolver, len(products))
	for i, product := range products {
		productsResolvers[i] = &ProductResolver{&product}
	}

	return productsResolvers, nil
}
