package in_memory_bd

import (
	"context"
	"time"

	"github.com/pkg/errors"
	entity "url-checker/internal/domain"
)

func (c *Cache) UpdateStatus(_ context.Context, url string, status entity.Status) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if val, ok := c.data[url]; ok {
		val.Status = int(status)
		val.LastCheck = time.Now().UnixMilli()
		c.data[url] = val

		return nil
	}

	return errors.New("Not fount url: " + url)
}
