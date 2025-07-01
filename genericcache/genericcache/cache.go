package genericcache

import (
	"slices"
	"sync"
	"time"
)

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time
}

// cache is a key value storage
type Cache[K comparable, V any] struct {
	// mu   sync.Mutex
	ttl               time.Duration
	mu                sync.Mutex
	data              map[K]entryWithTimeout[V]
	maxSize           int
	chronologicalKeys []K
}

// addKeyValue inserts a key and its value into the cache.
func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = entryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

// deleteKeyValue removes a key and its associated value from t
func (c *Cache[K, V]) deleteKeyValue(key K) {
	c.chronologicalKeys = slices.DeleteFunc(c.chronologicalKeys, func(k K) bool { return k == key })
	delete(c.data, key)
}

// New creates a usable Cache.
func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:               ttl,
		data:              make(map[K]entryWithTimeout[V]),
		maxSize:           maxSize,
		chronologicalKeys: make([]K, 0, maxSize),
	}
}

// Read returns the associated value for a key,
// and a boolean to true if the key is absent.

func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroV V
	e, ok := c.data[key]
	switch {
	case !ok:
		return zeroV, false
	case e.expires.Before(time.Now()):
		c.deleteKeyValue(key)
		return zeroV, false
	default:
		return e.value, true
	}
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, alreadyPresent := c.data[key]

	switch {
	case alreadyPresent:
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		c.deleteKeyValue(c.chronologicalKeys[0])
	}

	c.addKeyValue(key, value)
	return nil
}

// Delete removes the entry for the given key
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.deleteKeyValue(key)
}
