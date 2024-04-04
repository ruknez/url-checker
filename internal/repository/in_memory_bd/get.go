package in_memory_bd

import (
	"context"

	"url-checker/internal/repository/entity"
)

func (c *cache) Get(_ context.Context, url string) (entity.UrlInBd, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if data, ok := c.data[url]; ok {
		return data, nil
	}

	return entity.UrlInBd{}, entity.NoDataErr
}
