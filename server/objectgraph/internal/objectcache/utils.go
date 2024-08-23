package objectcache

import "fmt"

func (oc *ObjectCache) redisKey(groupKey string, objectKey string, valueKey string) string {
	return fmt.Sprintf("%s:%s:%s:%s", oc.prefix, groupKey, objectKey, valueKey)
}
