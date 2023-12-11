package cache

import (
	"errors"
	"sync"
)

type OrderCache struct {
	mutex    sync.RWMutex
	size     uint
	capacity uint
	orders   map[string][]byte
}

func NewOrderCache(capacity uint, orders map[string][]byte) *OrderCache {
	return &OrderCache{
		orders:   orders,
		size:     0,
		capacity: capacity,
	}
}

func (c *OrderCache) GetOrder(id string) ([]byte, bool) {
	c.mutex.RLock()

	defer c.mutex.RUnlock()

	order, hit := c.orders[id]
	if !hit {
		return nil, false
	}

	return order, true
}

func (c *OrderCache) AddOrder(id string, order []byte) (string, error) {

	c.mutex.Lock()

	defer c.mutex.Unlock()

	c.orders[id] = order

	return id, nil
}

func (c *OrderCache) DeleteOrder(id string) error {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	if _, hit := c.orders[id]; !hit {
		return errors.New("Cache: id not found")
	}

	delete(c.orders, id)

	return nil
}
