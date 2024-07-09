package retrievers

import (
	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

type RetrieveUser struct{}

func (*RetrieveUser) ByID(args objectgraph.RetrieverInput) (*database.User, error) {
	userID := args.PrimaryID.(uuid.UUID)

	user, err := database.SelectUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (*RetrieveUser) ByIDs(args objectgraph.RetrieverInput) ([]*database.User, error) {
	userIDs := args.PrimaryID.([]uuid.UUID)

	users, err := database.SelectUsersByIDs(userIDs)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (*RetrieveUser) Top1000() ([]*database.User, error) {
	users, err := database.SelectTop1000Users()
	if err != nil {
		return nil, err
	}

	return users, nil
}
