package retreivers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func GetProductByID(ctx context.Context, args engine.RetrieverArgs) (any, error) {
	productID, err := uuid.Parse(args["productID"])
	if err != nil {
		return nil, err
	}

	product, err := database.SelectProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return mapToProductResolver(product), nil
}

func GetProducts(ctx context.Context, args engine.RetrieverArgs) (any, error) {
	products, err := database.SelectProducts(ctx)
	if err != nil {
		return nil, err
	}

	r := make([]*resolvers.ProductResolver, len(products))
	for i, product := range products {
		r[i] = mapToProductResolver(product)
	}

	return &r, nil
}

func mapToProductResolver(product *database.Product) *resolvers.ProductResolver {
	if product == nil {
		return nil
	}

	return &resolvers.ProductResolver{
		Product: product,

		ProductID:                 scalars.NewUUID(product.ProductID),
		Name:                      product.Name,
		Volume:                    scalars.NullInt(product.Volume),
		ABV:                       scalars.NullInt(product.ABV),
		Price:                     scalars.NewPrice(product.Price),
		PromotionDiscount:         scalars.NewNilPrice(product.PromotionDiscount),
		PromotionDiscountMultiple: scalars.NullInt(product.PromotionDiscountMultiple),
		PromotionPrice:            scalars.NewPrice(calculatePromotionPrice(product.Price, product.PromotionDiscount)),
		UpdatedAt:                 scalars.NewNilTimestamp(product.UpdatedAt),
		CreatedAt:                 scalars.NewTimestamp(product.CreatedAt),
	}
}

func calculatePromotionPrice(price int, discount *int) int {
	value := price

	if discount != nil {
		value = price - *discount
	}

	return value
}
