package core

import (
	"github.com/nicholasguan/goacache/misc"
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex

	key   misc.CacheKeyType
	value interface{}

	liveDuration time.Duration
	modifyTime   time.Time
	createTime   time.Time
}

func NewCacheItem(key misc.CacheKeyType, value interface{}, liveDuration time.Duration) *CacheItem {
	return &CacheItem{
		key:          key,
		value:        value,
		liveDuration: liveDuration,
		createTime:   time.Now(),
		modifyTime:   time.Now(),
	}
}

func (item *CacheItem) GetKey() misc.CacheKeyType {
	return item.key
}

func (item *CacheItem) GetValue() interface{} {
	return item.value
}

func (item *CacheItem) SetValue(value interface{}) {
	item.value = value
	item.modifyTime = time.Now()
}

func (item *CacheItem) SetLiveDuration(liveDuration time.Duration) {
	item.liveDuration = liveDuration
	item.modifyTime = time.Now()
}
