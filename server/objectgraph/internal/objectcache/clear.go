package objectcache

import "context"

func (c *ObjectCache) Clear() error {
	for _, cacheGroup := range c.objectCache {
		cacheGroup.Flush()
	}

	_, err := c.redis.FlushAll(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}
