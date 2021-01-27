package core

import (
	"github.com/nicholasguan/goacache/misc"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex

	cacheMap map[misc.CacheKeyType]*CacheItem
}

func NewCacheTable() *CacheTable {
	return &CacheTable{
		cacheMap: make(map[misc.CacheKeyType]*CacheItem),
	}
}

func (table *CacheTable) AddItem(key misc.CacheKeyType, value interface{}, duration time.Duration) {
	table.Lock()
	defer table.Unlock()

	cacheItem := NewCacheItem(key, value, duration)
	table.cacheMap[key] = cacheItem
}

func (table *CacheTable) DelItem(key misc.CacheKeyType) {
	table.Lock()
	defer table.Unlock()

	delete(table.cacheMap, key)
}

func (table *CacheTable) Exists(key misc.CacheKeyType) bool {
	if !table.checkItemExpired(key) {
		return false
	}

	table.RLock()
	defer table.RUnlock()
	_, ok := table.cacheMap[key]

	return ok
}

func (table *CacheTable) SearchItem(key misc.CacheKeyType) *CacheItem {
	if !table.checkItemExpired(key) {
		return nil
	}

	table.RLock()
	defer table.RUnlock()
	item, ok := table.cacheMap[key]

	if !ok {
		return nil
	}

	return item
}

func (table *CacheTable) checkItemExpired(key misc.CacheKeyType) bool {
	table.RLock()
	item, ok := table.cacheMap[key]
	table.RUnlock()

	if !ok {
		return false
	}

	// 永不过期
	if item.liveDuration == 0 {
		return true
	}

	// 判断key是否过期，过期删掉
	if item.modifyTime.Add(item.liveDuration).Before(time.Now()) {
		table.Lock()
		delete(table.cacheMap, key)
		table.Unlock()
		return false
	}

	return true
}
