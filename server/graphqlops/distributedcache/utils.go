package distributedcache

func fmtKey(cacheKey string) string {
	return prefix + ":" + "distributedcache" + ":" + cacheKey
}
