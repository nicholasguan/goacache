package main

import (
	"fmt"
	goacache "github.com/nicholasguan/goacache/core"
	"github.com/nicholasguan/goacache/misc"
	"time"
)

func main() {
	table := goacache.NewCacheTable()

	var key, key1 misc.CacheKeyType = "test", "test1"
	table.AddItem(key, 123, 123*time.Second)

	fmt.Println(table.Exists(key))
	printItem(key, table.SearchItem(key))

	table.AddItem(key, 456, 123*time.Second)
	printItem(key, table.SearchItem(key))

	table.DelItem(key)
	printItem(key, table.SearchItem(key))

	fmt.Println(table.Exists(key1))
	printItem(key1, table.SearchItem(key1))
}

func printItem(key misc.CacheKeyType, item *goacache.CacheItem) {
	if item != nil {
		fmt.Printf("%s => %v %v\n", key, item.GetKey(), item.GetValue())
	} else {
		fmt.Printf("Can't find any item by key(%v)\n", key)
	}
}
