package graphql

import (
	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
)

var retreiversMap = engine.RetrieverMap{
	"get-brand":              getBrand,
	"get-products":           getProducts,
	"get-product":            getProduct,
	"get-product-categories": getProductCategories,
	"get-classification":     getClassification,
}

func getBrand(args engine.RetrieverArgs) (any, error) {
	brandID, err := uuid.Parse(args["brandID"])
	if err != nil {
		return nil, err
	}

	brand, err := database.SelectBrandByID(brandID)
	if err != nil {
		return nil, err
	}

	if brand == nil {
		return nil, nil
	}

	r := &resolvers.BrandResolver{Brand: brand}

	return r, nil
}

func getProduct(args engine.RetrieverArgs) (any, error) {
	productID, err := uuid.Parse(args["productID"])
	if err != nil {
		return nil, err
	}

	product, err := database.SelectProductByID(productID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	r := &resolvers.ProductResolver{Product: product}

	return r, nil
}

func getClassification(args engine.RetrieverArgs) (any, error) {
	classificationID, err := uuid.Parse(args["classificationID"])
	if err != nil {
		return nil, err
	}

	classification, err := database.SelectClassificationByID(classificationID)
	if err != nil {
		return nil, err
	}

	if classification == nil {
		return nil, nil
	}

	r := &resolvers.ClassificationResolver{Classification: classification}

	return r, nil
}

func getProducts(args engine.RetrieverArgs) (any, error) {
	products, err := database.SelectProducts()
	if err != nil {
		return nil, err
	}

	r := make([]*resolvers.ProductResolver, len(products))
	for i, product := range products {
		r[i] = &resolvers.ProductResolver{Product: product}
	}

	return &r, nil
}

func getProductCategories(args engine.RetrieverArgs) (any, error) {
	productID, err := uuid.Parse(args["productID"])
	if err != nil {
		return nil, err
	}

	categories, err := database.SelectCategoriesByProductID(productID)
	if err != nil {
		return nil, err
	}

	r := make([]*resolvers.CategoryResolver, 0, len(categories))

	for i := range categories {
		r = append(r, &resolvers.CategoryResolver{Category: categories[i]})
	}

	return &r, nil
}
