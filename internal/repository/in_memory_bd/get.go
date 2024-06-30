package in_memory_bd

import (
	"context"

	entity "url-checker/internal/domain"
	"url-checker/internal/repository/in_memory_bd/mapping"
)

func (c *Cache) Get(_ context.Context, url string) (entity.UrlInfo, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	if data, ok := c.data[url]; ok {
		return mapping.URLBdToURLInfoMapping(data), nil
	}

	return entity.UrlInfo{}, entity.NoDataErr
}

func (c *Cache) GetAllUrls(_ context.Context) []string {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	res := make([]string, 0, len(c.data))
	for key := range c.data {
		res = append(res, key)
	}

	return res
}
