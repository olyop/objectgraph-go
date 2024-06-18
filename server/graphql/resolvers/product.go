package resolvers

import (
	"context"

	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
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
	return graphqlops.ResolverList[CategoryResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveProductCategories",
		RetrieverArgs: graphqlops.RetrieverArgs{"productID": r.Product.ProductID},
	})
}

func (r *ProductResolver) Brand(ctx context.Context) (*BrandResolver, error) {
	return graphqlops.Resolver[BrandResolver](ctx, graphqlops.ResolverOptions{
		CacheDuration: "catalog",
		RetrieverKey:  "RetrieveBrandByID",
		RetrieverArgs: graphqlops.RetrieverArgs{"brandID": r.Product.BrandID},
	})
}
