package objectcache

import "context"

func (c *ObjectCache) Clear() error {
	err := c.redis.FlushAll(context.Background()).Err()
	if err != nil {
		return err
	}

	return nil
}
