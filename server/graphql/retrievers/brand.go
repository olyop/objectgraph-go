package retrievers

import (
	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

type RetrieveBrand struct{}

func (*RetrieveBrand) ByID(args objectgraph.RetrieverInput) (*database.Brand, error) {
	brandID := args.PrimaryID.(uuid.UUID)

	brand, err := database.SelectBrandByID(brandID)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func (*RetrieveBrand) ByIDs(args objectgraph.RetrieverInput) ([]*database.Brand, error) {
	brandIDs := args.PrimaryID.([]uuid.UUID)

	brands, err := database.SelectBrandsByIDs(brandIDs)
	if err != nil {
		return nil, err
	}

	return brands, nil
}
