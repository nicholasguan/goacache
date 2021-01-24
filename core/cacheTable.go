package core

import (
	"sync"
)

type CacheTable struct {
	sync.RWMutex

	cacheMap map[string]*CacheItem
}

func NewCacheTable() *CacheTable {
	return &CacheTable{
		cacheMap: make(map[string]*CacheItem),
	}
}

func (table *CacheTable) AddItem(key string, value interface{}) {
	table.Lock()
	defer table.Unlock()

	cacheItem := NewCacheItem(key, value)
	table.cacheMap[key] = cacheItem
}

func (table *CacheTable) DelItem(key string) {
	table.Lock()
	defer table.Unlock()

	delete(table.cacheMap, key)
}

func (table *CacheTable) Exists(key string) bool {
	table.RLock()
	defer table.RUnlock()
	_, ok := table.cacheMap[key]

	return ok
}

func (table *CacheTable) SearchItem(key string) *CacheItem {
	table.RLock()
	defer table.RUnlock()
	item, ok := table.cacheMap[key]

	if !ok {
		return nil
	}

	return item
}
