package inmemorycache

func Get[T any](groupKey string, cacheKey string) (*T, bool) {
	groupCache := handleGroup(groupKey)

	item, exists := groupCache.Get(cacheKey)
	if !exists {
		return nil, false
	}

	value := item.(*T)

	return value, true
}
