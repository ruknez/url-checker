package in_memory_bd

import (
	"context"
	"time"
)

func (c *Cache) GetAllUrlsToCheck(_ context.Context) []string {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	res := make([]string, 0, len(c.data))
	for key, val := range c.data {
		if time.Now().UnixMilli() <= val.LastCheck+val.Duration {
			res = append(res, key)
		}
	}

	return res
}
