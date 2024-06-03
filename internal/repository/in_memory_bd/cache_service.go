package in_memory_bd

import (
	"sync"

	inMemoryBd "url-checker/internal/repository/in_memory_bd/entity"
)

type Cache struct {
	data map[string]inMemoryBd.UrlInBd
	mtx  sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]inMemoryBd.UrlInBd, 10),
		mtx:  sync.RWMutex{},
	}
}
