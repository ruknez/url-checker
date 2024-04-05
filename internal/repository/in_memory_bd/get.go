package in_memory_bd

import (
	"context"

	"url-checker/internal/repository/entity"
)

func (c *cache) Get(_ context.Context, url string) (entity.UrlInBd, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	if data, ok := c.data[url]; ok {
		return data, nil
	}

	return entity.UrlInBd{}, entity.NoDataErr
}

func (c *cache) GetAllUrls(_ context.Context) []string {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	res := make([]string, 0, len(c.data))
	for key := range c.data {
		res = append(res, key)
	}

	return res
}
