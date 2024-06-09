package inmemorycache

func Get[R any](groupKey string, cacheKey string) (R, bool) {
	groupCache := handleGroup(groupKey)

	var value R

	item, exists := groupCache.Get(cacheKey)
	if !exists {
		return value, false
	}

	value = item.(R)

	return value, true
}
