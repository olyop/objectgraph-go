package resolvers

import (
	"errors"

	"github.com/olyop/graphql-go/database"
	"github.com/olyop/graphql-go/resolvers/scalars"
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

func (r *ProductResolver) CreatedAt() scalars.Timestamp {
	return scalars.Timestamp{Time: r.Product.CreatedAt}
}

func (r *ProductResolver) Price() (scalar scalars.Price, err error) {
	return scalars.Price{Value: r.Product.Price}, nil
}

func (r *ProductResolver) Volume() *int32 {
	value := int32(r.Product.Volume)

	return &value
}

func (r *ProductResolver) ABV() *int32 {
	value := int32(r.Product.ABV)

	return &value
}

func (r *ProductResolver) Brand() (*BrandResolver, error) {
	brand, err := database.SelectBrandByID(r.Product.BrandID)
	if err != nil {
		return nil, errors.New("failed to get brand")
	}

	return &BrandResolver{&brand}, nil
}

func (r *ProductResolver) Categories() ([]*CategoryResolver, error) {
}
