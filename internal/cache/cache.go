package cache

import "sync"

// Cache provides a simple cache of expensive resources. This is intended for a
// narrow scope, where a well-known set of resources can be cached. The cache
// has no eviction policy, so any cached resource will be kept indefinitely.
//
// One example use case is to cached parsed JavaScript that is installed as
// polyfill in the browser.
type Cache struct {
	data map[string]any
	mu   sync.RWMutex
}

// Get returns a resource from the cache found by key, assignable to type T. If
// key is not found, or the type is not T, ok will be false. Panics if c is nil.
func Get[T any](c *Cache, key string) (res T, ok bool) {
	if c == nil {
		panic("gost-dom/cache: Get called with nil cache")
	}
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

// Set adds a to the cache with the specified key. If the key already exists,
// the value is overwritten. Panics if c is nil.
func Set[T any](c *Cache, key string, val T) {
	if c == nil {
		panic("gost-dom/cache: Set called with nil cache")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.data == nil {
		c.data = make(map[string]any)
	}
	c.data[key] = val
}
