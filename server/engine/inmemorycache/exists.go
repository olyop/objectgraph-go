package inmemorycache

func Exists(groupKey string, cacheKey string) bool {
	groupCache := handleGroup(groupKey)

	_, exists := groupCache.Get(cacheKey)

	return exists
}
