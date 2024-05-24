package store

import (
	"time"

	"github.com/google/uuid"
	"github.com/olyop/graphql-go/server/database"
	"github.com/olyop/graphql-go/server/store/cache"
)

// check cache
// if cache is expired
//   re-fetch
// else
//   serve cache

func BrandStore[T any](brandID uuid.UUID, key string) (T, error) {
	id := brandID.String()

	var value T

	cached, found := cache.Get[map[string]any]("brands", id)
	if found {
		value = cached[key].(T)

		return value, nil
	}

	data, err := brandRetriever(brandID)
	if err != nil {
		return value, err
	}

	value = data[key].(T)

	cache.Set("brands", id, data, time.Second*15)

	return value, nil
}

func brandRetriever(brandID uuid.UUID) (map[string]any, error) {
	brand, err := database.SelectBrandByID(brandID)
	if err != nil {
		return nil, err
	}

	m := map[string]any{
		"brandID":   brand.BrandID,
		"name":      brand.Name,
		"updatedAt": brand.UpdatedAt,
		"createdAt": brand.CreatedAt,
	}

	return m, nil
}
