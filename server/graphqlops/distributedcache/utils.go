package distributedcache

func fmtKey(cacheKey string) string {
	return keyPrefix + ":" + "cache" + ":" + cacheKey
}
