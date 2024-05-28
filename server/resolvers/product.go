package resolvers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/resolvers/scalars"
)

type ProductResolver struct {
	*database.Product
}

func (r *ProductResolver) ProductID() scalars.UUID {
	return scalars.NewUUID(r.Product.ProductID)
}

func (r *ProductResolver) Name() string {
	return r.Product.Name
}

func (r *ProductResolver) UpdatedAt() *scalars.Timestamp {
	return scalars.NewNillTimestamp(r.Product.UpdatedAt)
}

func (r *ProductResolver) CreatedAt() scalars.Timestamp {
	return scalars.NewTimestamp(r.Product.CreatedAt)
}

func (r *ProductResolver) Price() scalars.Price {
	return scalars.NewPrice(r.Product.Price)
}

func (r *ProductResolver) PromotionDiscount() *scalars.Price {
	return scalars.NewNillPrice(r.Product.PromotionDiscount)
}

func (r *ProductResolver) PromotionDiscountMultiple() *int32 {
	return handleSqlNullInt32(r.Product.PromotionDiscountMultiple)
}

func (r *ProductResolver) Volume() *int32 {
	return handleSqlNullInt32(r.Product.Volume)
}

func (r *ProductResolver) ABV() *float64 {
	return handleSqlNullFloat64(r.Product.ABV)
}

func (r *ProductResolver) Categories() ([]*CategoryResolver, error) {
	return engine.Resolver(engine.ResolverOptions[[]*CategoryResolver]{
		GroupKey: "product-categories",
		Duration: 15 * time.Second,
		CacheKey: r.Product.ProductID.String(),
		Retrieve: categoriesRetriever(r.Product.ProductID),
	})
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

func (r *ProductResolver) Brand(ctx context.Context) (*BrandResolver, error) {
	return engine.Resolver(engine.ResolverOptions[*BrandResolver]{
		GroupKey: "product-brands",
		Duration: 15 * time.Second,
		CacheKey: r.Product.BrandID.String(),
		Retrieve: brandRetriever(r.Product.BrandID),
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
