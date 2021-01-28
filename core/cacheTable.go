package core

import (
	"fmt"
	"github.com/nicholasguan/goacache/misc"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex

	cacheMap map[misc.CacheKeyType]*CacheItem

	checkExpiredTicker *time.Ticker
}

func NewCacheTable() *CacheTable {
	checkTicker := time.NewTicker(time.Second)
	cacheTable := &CacheTable{
		cacheMap:           make(map[misc.CacheKeyType]*CacheItem),
		checkExpiredTicker: checkTicker,
	}

	go func() {
		for {
			select {
			case <-checkTicker.C:
				fmt.Println("Self-checking is doing")
				cacheTable.checkAllItemExpired()
			}
		}
	}()

	return cacheTable
}

func (table *CacheTable) AddItem(key misc.CacheKeyType, value interface{}, duration time.Duration) {
	table.Lock()
	defer table.Unlock()

	cacheItem := NewCacheItem(key, value, duration)
	table.cacheMap[key] = cacheItem
	fmt.Println("add item", key, value)
}

func (table *CacheTable) DelItem(key misc.CacheKeyType) {
	table.Lock()
	defer table.Unlock()

	delete(table.cacheMap, key)
	fmt.Println("delete item", key)
}

func (table *CacheTable) Exists(key misc.CacheKeyType) bool {
	table.RLock()
	defer table.RUnlock()
	_, ok := table.cacheMap[key]

	return ok
}

func (table *CacheTable) SearchItem(key misc.CacheKeyType) *CacheItem {
	table.RLock()
	defer table.RUnlock()
	item, ok := table.cacheMap[key]

	if !ok {
		return nil
	}

	return item
}

func (table *CacheTable) checkAllItemExpired() {
	for key := range table.cacheMap {
		table.checkItemExpired(key)
	}
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
