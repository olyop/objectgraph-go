package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

type ProductResolver struct {
	Product *database.Product

	ProductID                 scalars.UUID
	Name                      string
	Volume                    *int32
	ABV                       *int32
	Price                     scalars.Price
	PromotionDiscount         *scalars.Price
	PromotionDiscountMultiple *int32
	PromotionPrice            scalars.Price
	UpdatedAt                 *scalars.Timestamp
	CreatedAt                 scalars.Timestamp
}

func (r *ProductResolver) Categories(ctx context.Context) ([]*CategoryResolver, error) {
	return engine.ResolverList[CategoryResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "retrieve-product-categories",
		RetrieverArgs: engine.RetrieverArgs{"productID": r.Product.ProductID.String()},
	})
}

func (r *ProductResolver) Brand(ctx context.Context) (*BrandResolver, error) {
	return engine.Resolver[BrandResolver](ctx, engine.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "retrieve-brand-by-id",
		RetrieverArgs: engine.RetrieverArgs{"brandID": r.Product.BrandID.String()},
	})
}
