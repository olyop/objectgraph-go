package objectcache

import "context"

func (c *ObjectCache) Clear(ctx context.Context) error {
	c.clearLock.Lock()
	defer c.clearLock.Unlock()

	for _, cacheGroup := range c.cacheGroups {
		cacheGroup.Flush()
	}

	_, err := c.redis.FlushAll(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
