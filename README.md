# ObjectGraph Go

Repo for development of a GraphQL framework currently named 'ObjectGraph'.

The goal is to create a framework that simplifies the caching & data fetching layer in the context of a GraphQL server.
It should be as easy as defining a few 'Retreivers' for each data model then using schema annotations to link the data fields and retreivers together any way the developer wants.
The ObjectGraph engine will handle linking everything together and resolving queries and caching (authentication in the future).
It should not complicate the layer and should work in a predictable schema-first way while also being performant and scalable aimed for serverless invocations.

```graphql
type Query {
	getProductByID(productID: UUID!): Product!
		@retrieve(key: "Product/ByID", args: ["primaryID=$productID"])
}

type Product {
	productID: UUID!
		@object(key: "Product/ProductID")
		@cache(ttl: "1d")

	name: String!
		@object(key: "Product/Name")
		@cache(ttl: "1d")

	price: Int!
		@object(key: "Product/Price")
		@cache(ttl: "1d")

	promotionDiscount: Int
		@object(key: "Product/PromotionDiscount")
		@cache(ttl: "1d")

	promotionDiscountMultiple: Int
		@object(key: "Product/PromotionDiscountMultiple")
		@cache(ttl: "1d")

	volume: Int
		@object(key: "Product/Volume")
		@cache(ttl: "1d")

	abv: Int
		@object(key: "Product/ABV")
		@cache(ttl: "1d")

	updatedAt: Timestamp
		@object(key: "Product/UpdatedAt")
		@cache(ttl: "1d")

	createdAt: Timestamp!
		@object(key: "Product/CreatedAt")
		@cache(ttl: "1d")


	brand: Brand!
		@retrieve(key: "Brand/ByID", args: ["primaryID=Product/BrandID"])
		@cache(ttl: "1d")

	categories: [Category!]!
		@retrieve(key: "Category/AllByProductID", args: ["productID=Product/ProductID"])
		@cache(ttl: "1d")
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

For a full example see the [server/graphql/schema](https://github.com/olyop/objectgraph-go/tree/main/server/graphql/schema) directory.