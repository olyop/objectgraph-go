package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers/scalars"
)

type ProductResolver struct {
	Product *database.Product
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

func (r *ProductResolver) Categories(ctx context.Context) ([]*CategoryResolver, error) {
	return engine.ResolverList[CategoryResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-product-categories",
		RetrieverArgs: engine.RetrieverArgs{"productID": r.Product.ProductID.String()},
	})
}

func (r *ProductResolver) Brand(ctx context.Context) (*BrandResolver, error) {
	return engine.Resolver[BrandResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "get-brand",
		RetrieverArgs: engine.RetrieverArgs{"brandID": r.Product.BrandID.String()},
	})
}
