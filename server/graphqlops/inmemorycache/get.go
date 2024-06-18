package inmemorycache

func Get(groupKey string, cacheKey string) (any, bool) {
	groupCache := handleGroup(groupKey)

	item, exists := groupCache.Get(cacheKey)
	if !exists {
		return nil, false
	}

	return item, true
}
