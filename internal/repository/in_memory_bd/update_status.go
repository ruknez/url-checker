package in_memory_bd

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func (c *cache) UpdateStatus(_ context.Context, url string, status int) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if val, ok := c.data[url]; ok {
		val.Status = status
		val.LastCheck = time.Now().UnixMilli()
		c.data[url] = val

		return nil
	}

	return errors.New("Not fount url: " + url)
}
