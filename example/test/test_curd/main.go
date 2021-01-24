package main

import (
	"fmt"
	goacache "github.com/nicholasguan/goacache/core"
)

func main() {
	table := goacache.NewCacheTable()

	key, key1 := "test", "test1"
	table.AddItem(key, 123)

	fmt.Println(table.Exists(key))
	printItem(key, table.SearchItem(key))

	table.AddItem(key, 456)
	printItem(key, table.SearchItem(key))

	table.DelItem(key)
	printItem(key, table.SearchItem(key))

	fmt.Println(table.Exists(key1))
	printItem(key1, table.SearchItem(key1))
}

func printItem(key string, item *goacache.CacheItem) {
	if item != nil {
		fmt.Printf("%s => %v %v\n", key, item.GetKey(), item.GetValue())
	} else {
		fmt.Printf("Can't find any item by key(%v)\n", key)
	}
}
