package inmemorycache

func Delete(groupKey string, cacheKey string) {
	groupCache := handleGroup(groupKey)

	groupCache.Delete(cacheKey)
}
