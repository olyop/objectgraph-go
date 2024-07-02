package retrievers

import (
	"github.com/google/uuid"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph"
)

type RetrieveContact struct{}

func (*RetrieveContact) AllByUserID(args objectgraph.RetrieverArgs) ([]*database.Contact, error) {
	userID := args.GetPrimary().(uuid.UUID)

	contacts, err := database.SelectContactsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}
