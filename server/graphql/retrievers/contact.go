package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/graphql/enums"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
	"github.com/olyop/graphql-go/server/graphqlops"
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

	return r, nil
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
