package cache

import "sync"

type Cache struct {
	data map[string]any
	mu   sync.RWMutex
}

func Get[T any](c *Cache, key string) (res T, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.data == nil {
		ok = false
		return
	}
	var val any
	val, ok = c.data[key]
	if !ok {
		return
	}
	res, ok = val.(T)
	return
}

func Set[T any](c *Cache, key string, val T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.data == nil {
		c.data = make(map[string]any)
	}
	c.data[key] = val
}
