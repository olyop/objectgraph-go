package retrievers

import (
	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

type RetrieveCategory struct{}

func (*RetrieveCategory) ByID(args objectgraph.RetrieverArgs) (*database.Category, error) {
	categoryID := args.GetPrimary().(uuid.UUID)

	category, err := database.SelectCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (*RetrieveCategory) ByIDs(args objectgraph.RetrieverArgs) ([]*database.Category, error) {
	categoryIDs := args.GetPrimary().([]uuid.UUID)

	categories, err := database.SelectCategoriesByIDs(categoryIDs)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (*RetrieveCategory) AllByProductID(args objectgraph.RetrieverArgs) ([]*database.Category, error) {
	productID := args.GetArg("productID").(uuid.UUID)

	categories, err := database.SelectCategoriesByProductID(productID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
