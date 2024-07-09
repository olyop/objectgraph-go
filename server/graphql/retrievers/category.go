package retrievers

import (
	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

type RetrieveCategory struct{}

func (*RetrieveCategory) ByID(input objectgraph.RetrieverInput) (*database.Category, error) {
	categoryID := input.PrimaryID.(uuid.UUID)

	category, err := database.SelectCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (*RetrieveCategory) ByIDs(input objectgraph.RetrieverInput) ([]*database.Category, error) {
	categoryIDs := input.PrimaryID.([]uuid.UUID)

	categories, err := database.SelectCategoriesByIDs(categoryIDs)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (*RetrieveCategory) AllByProductID(input objectgraph.RetrieverInput) ([]*database.Category, error) {
	productID := input.Args["productID"].(uuid.UUID)

	categories, err := database.SelectCategoriesByProductID(productID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
