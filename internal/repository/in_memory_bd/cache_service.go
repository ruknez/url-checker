package in_memory_bd

import (
	"sync"

	inMemoryBd "url-checker/internal/repository/in_memory_bd/entity"
)

type cache struct {
	data map[string]inMemoryBd.UrlInBd
	mtx  sync.RWMutex
}

func NewCache() *cache {
	return &cache{
		data: make(map[string]inMemoryBd.UrlInBd, 10),
		mtx:  sync.RWMutex{},
	}
}
