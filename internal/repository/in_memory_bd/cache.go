package in_memory_bd

import (
	"sync"

	"url-checker/internal/repository/entity"
)

type cache struct {
	data map[string]entity.UrlInBd
	mtx  sync.Mutex
}

func NewCache() *cache {
	return &cache{
		data: make(map[string]entity.UrlInBd, 10),
		mtx:  sync.Mutex{},
	}
}
