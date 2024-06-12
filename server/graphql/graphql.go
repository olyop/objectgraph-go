package graphql

import (
	"time"

	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/retrievers"
)

func Initialize() {
	engine.RegisterRetrievers(engine.RetrieverMap{
		"retrieve-brand-by-id":          retrievers.RetreiveBrandByID,
		"retrieve-classification-by-id": retrievers.RetreiveClassificationByID,
		"retrieve-product-by-id":        retrievers.RetreiveProductByID,
		"retrieve-product-categories":   retrievers.RetreiveProductCategories,
		"retrieve-top-1000-products":    retrievers.RetreiveTop1000Products,
		"retrieve-top-1000-users":       retrievers.RetreiveTop1000Users,
		"retrieve-user-contacts":        retrievers.RetreiveUserContacts,
	})

	engine.RegisterCacheDurations(map[string]time.Duration{
		"catalog": 45 * time.Second,
	})
}
