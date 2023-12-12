package main

import (
	"container/list"
	"fmt"
	"sync"
	"unsafe"

	"github.com/golang/groupcache/lru"
)

type OrderCache struct {
	sync.RWMutex
	size     uint
	capacity uint
	orders   map[string][]byte
	queue    *[]list.Element
}

func main() {
	var cache OrderCache
	fmt.Println(unsafe.Sizeof(cache))

	cached := lru.New(100)

	cached.RemoveOldest()

}
