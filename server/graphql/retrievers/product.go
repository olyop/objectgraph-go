package retrievers

import (
	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

type RetrieveProduct struct{}

func (*RetrieveProduct) ByID(args objectgraph.RetrieverArgs) (*database.Product, error) {
	productID := args.GetPrimary().(uuid.UUID)

	product, err := database.SelectProductByID(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (*RetrieveProduct) Top1000(args objectgraph.RetrieverArgs) ([]*database.Product, error) {
	products, err := database.SelectTop1000Products()
	if err != nil {
		return nil, err
	}

	return products, nil
}
