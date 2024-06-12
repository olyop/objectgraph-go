package retrievers

import (
	"context"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
	"github.com/olyop/graphql-go/server/graphql/scalars"
)

func RetreiveUserContacts(ctx context.Context, args engine.RetrieverArgs) (any, error) {
	userID, err := uuid.Parse(args["userID"])
	if err != nil {
		return nil, err
	}

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
	if contact == nil {
		return nil
	}

	return &resolvers.ContactResolver{
		ContactID: scalars.NewUUID(contact.ContactID),
		Value:     contact.Value,
		Type:      contact.Type,
		UpdatedAt: scalars.NewNilTimestamp(contact.UpdatedAt),
		CreatedAt: scalars.NewTimestamp(contact.CreatedAt),
	}
}
