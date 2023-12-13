package cache

import (
	"container/list"
	"fmt"
	"sync"
)

type Item struct {
	Key   string
	Value interface{}
}

type OrderCache struct {
	sync.RWMutex
	capacity int
	orders   map[string]*list.Element
	queue    *list.List
}

func NewOrderCache(capacity int) *OrderCache {
	return &OrderCache{
		orders:   make(map[string]*list.Element, capacity),
		capacity: capacity,
		queue:    list.New(),
	}
}

func (c *OrderCache) GetCapacity() int {
	return c.capacity
}

func (c *OrderCache) GetOrderById(id string) ([]byte, bool) {
	c.RLock()

	defer c.RUnlock()

	element, exists := c.orders[id]
	if !exists {
		return nil, false
	}

	c.queue.MoveToFront(element)

	//c.printCacheData()

	return element.Value.(*Item).Value.([]byte), true
}

func (c *OrderCache) AddOrder(id string, order []byte) (string, error) {
	c.Lock()

	defer c.Unlock()

	if element, exists := c.orders[id]; exists {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = order
		return id, fmt.Errorf("repository.cache.AddOrder: order with id = %s already exists", id)
	}

	if c.queue.Len() == c.capacity {
		c.purgeOrder()
	}

	item := &Item{
		Key:   id,
		Value: order,
	}

	element := c.queue.PushFront(item)
	c.orders[item.Key] = element

	return id, nil
}

func (c *OrderCache) DeleteOrder(id string) error {
	c.Lock()

	defer c.Unlock()

	if _, hit := c.orders[id]; !hit {
		return fmt.Errorf("repository.cache.DeleteOrder: id not found")
	}

	delete(c.orders, id)

	return nil
}

func (c *OrderCache) purgeOrder() {
	if element := c.queue.Back(); element != nil {
		id := c.queue.Remove(element).(*Item).Key
		delete(c.orders, id)
	}
}

func (c *OrderCache) printCacheData() {
	elem := c.queue.Front()
	fmt.Println(elem.Value.(*Item).Key)
	for ; elem.Next() != nil; elem = elem.Next() {
		fmt.Println(elem.Next().Value.(*Item).Key)
	}
}
