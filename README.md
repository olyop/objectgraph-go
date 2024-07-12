# ObjectGraph Go

Repo for development of a GraphQL framework currently named 'ObjectGraph'.

The goal is to create a framework that simplifies the caching & data fetching layer in the context of a GraphQL server.
It should be as easy as defining a few 'Retreivers' for each data model then using schema annotations to link the data fields and retreivers together any way the developer wants.
The ObjectGraph engine will handle linking everything together and resolving queries and caching (authentication in the future).
It should not complicate the layer and should work in a predictable schema-first way while also being performant and scalable aimed for serverless invocations.

```graphql
type Query {
	getTop1000Users: [User!]! @retrieve(key: "User/Top1000")
	getUserByID(userID: UUID!): User! @retrieve(key: "User/ByID", args: ["primaryID=$userID"])
}

type User {
	userID: UUID! @object(field: "User/UserID") @cache(ttl: "1d")
	userName: String! @object(field: "User/UserName") @cache(ttl: "1d")
	firstName: String! @object(field: "User/FirstName") @cache(ttl: "1d")
	lastName: String! @object(field: "User/LastName") @cache(ttl: "1d")
	dob: Timestamp! @object(field: "User/DOB") @cache(ttl: "1d")
	updatedAt: Timestamp @object(field: "User/UpdatedAt") @cache(ttl: "1d")
	createdAt: Timestamp! @object(field: "User/CreatedAt") @cache(ttl: "1d")

	contacts: [Contact!]! @retrieve(key: "Contact/AllByUserID", args: ["userID=User/UserID"])
}
```

```go
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
```
