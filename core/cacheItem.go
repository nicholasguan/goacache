package core

import (
	"sync"
)

type CacheItem struct {
	sync.RWMutex

	key   string
	value interface{}
}

func NewCacheItem(key string, value interface{}) *CacheItem {
	return &CacheItem{
		key:   key,
		value: value,
	}
}

func (item *CacheItem) GetKey() string {
	return item.key
}

func (item *CacheItem) GetValue() interface{} {
	return item.value
}
