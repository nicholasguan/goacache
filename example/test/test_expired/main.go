package main

import (
	"fmt"
	goacache "github.com/nicholasguan/goacache/core"
	"github.com/nicholasguan/goacache/misc"
	"time"
)

func main() {
	table := goacache.NewCacheTable()

	duration := 4 * time.Second

	var key, key1 misc.CacheKeyType = "test", "test1"
	table.AddItem(key, 123, duration)

	// test key
	fmt.Println(table.Exists(key))
	printItem(key, table.SearchItem(key))

	// test timeOut
	currentTime := time.Now()
	for table.Exists(key) {
		fmt.Printf("第%d秒：\n", time.Now().Sub(currentTime)/time.Second)
		printItem(key, table.SearchItem(key))
		time.Sleep(time.Second)
	}
	printItem(key, table.SearchItem(key))

	// test reset timeout key1
	table.AddItem(key1, 456, duration)
	fmt.Println(table.Exists(key1))
	printItem(key1, table.SearchItem(key1))

	// test timeOut
	currentTime = time.Now()
	for table.Exists(key1) {
		currentSecond := time.Now().Sub(currentTime) / time.Second
		fmt.Printf("第%d秒：\n", currentSecond)
		if currentSecond == duration/time.Second/2 {
			// 直接新建
			// table.AddItem(key1, "修改value", duration)
			//
			// 修改value
			//    fmt.Println("change value")
			//    item := table.SearchItem(key1)
			//    item.SetValue("asd")
			// 修改liveDuration

			item := table.SearchItem(key1)
			item.SetLiveDuration(2 * duration)
		}
		printItem(key1, table.SearchItem(key1))

		time.Sleep(time.Second)
	}
	printItem(key1, table.SearchItem(key1))
}

func printItem(key misc.CacheKeyType, item *goacache.CacheItem) {
	if item != nil {
		fmt.Printf("%s => %+v\n", key, item)
	} else {
		fmt.Printf("Can't find any item by key(%v)\n", key)
	}
}
