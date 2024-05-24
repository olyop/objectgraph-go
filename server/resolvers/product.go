package resolvers

import (
	"time"

	"github.com/google/uuid"
	// "github.com/graph-gophers/dataloader"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/resolvers/scalars"
)

type ProductResolver struct {
	*database.Product
}

func (r *ProductResolver) ProductID() scalars.UUID {
	return scalars.UUID{UUID: r.Product.ProductID}
}

func (r *ProductResolver) Name() string {
	return r.Product.Name
}

func (r *ProductResolver) UpdatedAt() *scalars.Timestamp {
	if r.Product.UpdatedAt.IsZero() {
		return nil
	}

	return &scalars.Timestamp{Time: r.Product.UpdatedAt}
}

func (r *ProductResolver) CreatedAt() scalars.Timestamp {
	return scalars.Timestamp{Time: r.Product.CreatedAt}
}

func (r *ProductResolver) Price() scalars.Price {
	return scalars.Price{Value: r.Product.Price}
}

func (r *ProductResolver) Volume() *int32 {
	if !r.Product.Volume.Valid {
		return nil
	}

	value := int32(r.Product.Volume.Int64)

	return &value
}

func (r *ProductResolver) ABV() *float64 {
	if !r.Product.ABV.Valid {
		return nil
	}

	value := float64(r.Product.ABV.Float64)

	return &value
}

func (r *ProductResolver) Brand() (*BrandResolver, error) {
	return engine.Resolver(engine.ResolverOptions[BrandResolver]{
		GroupKey: "brands",
		Duration: 15 * time.Second,
		CacheKey: r.Product.BrandID.String(),
		Retrieve: brandRetriever(r.Product.BrandID),
	})
}

func (r *ProductResolver) Categories() ([]*CategoryResolver, error) {
	return engine.Resolvers(engine.ResolversOptions[CategoryResolver]{
		GroupKey: "categories",
		Duration: 15 * time.Second,
		CacheKey: r.Product.ProductID.String(),
		Retrieve: categoriesRetriever(r.Product.ProductID),
	})
}

func brandRetriever(brandID uuid.UUID) func() (*BrandResolver, error) {
	return func() (*BrandResolver, error) {
		brand, err := database.SelectBrandByID(brandID)
		if err != nil {
			return nil, err
		}

		return &BrandResolver{brand}, nil
	}
}

func categoriesRetriever(productID uuid.UUID) func() ([]*CategoryResolver, error) {
	return func() ([]*CategoryResolver, error) {
		categories, err := database.SelectCategoriesByProductID(productID)
		if err != nil {
			return nil, err
		}

		resolvers := make([]*CategoryResolver, 0, len(categories))

		for i := range categories {
			resolvers = append(resolvers, &CategoryResolver{categories[i]})
		}

		return resolvers, nil
	}
}
