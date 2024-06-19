package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphqlops-go/database"
	"github.com/olyop/graphqlops-go/graphql/enums"
	"github.com/olyop/graphqlops-go/graphql/resolvers"
	"github.com/olyop/graphqlops-go/graphql/scalars"
	"github.com/olyop/graphqlops-go/graphqlops"
)

func (*Retrievers) RetrieveUserContacts(ctx context.Context, args graphqlops.RetrieverArgs) (any, error) {
	userID := args["userID"].(uuid.UUID)

	contacts, err := database.SelectContactsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	r := make([]*resolvers.ContactResolver, len(contacts))
	for i := range contacts {
		r[i] = mapToContactResolver(contacts[i])
	}

	return &r, nil
}

func mapToContactResolver(contact *database.Contact) *resolvers.ContactResolver {
	return &resolvers.ContactResolver{
		ContactID:   scalars.NewUUID(contact.ContactID),
		Value:       contact.Value,
		ContactType: enums.NewContactType(contact.Type),
		UpdatedAt:   scalars.NewNilTimestamp(contact.UpdatedAt),
		CreatedAt:   scalars.NewTimestamp(contact.CreatedAt),
	}
}
