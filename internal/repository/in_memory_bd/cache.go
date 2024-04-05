package in_memory_bd

import (
	"sync"

	"url-checker/internal/repository/entity"
)

type cache struct {
	data map[string]entity.UrlInBd
	mtx  sync.RWMutex
}

func NewCache() *cache {
	return &cache{
		data: make(map[string]entity.UrlInBd, 10),
		mtx:  sync.RWMutex{},
	}
}
