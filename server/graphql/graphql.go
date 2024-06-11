package graphql

import (
	"time"

	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/retreivers"
)

func Initialize() {
	engine.RegisterRetrievers(engine.RetrieverMap{
		"get-brand-by-id":          retreivers.GetBrandByID,
		"get-products":             retreivers.GetProducts,
		"get-product-by-id":        retreivers.GetProductByID,
		"get-product-categories":   retreivers.GetProductCategories,
		"get-classification-by-id": retreivers.GetClassificationByID,
	})

	engine.RegisterCacheDurations(map[string]time.Duration{
		"catalog": 45 * time.Second,
	})
}
