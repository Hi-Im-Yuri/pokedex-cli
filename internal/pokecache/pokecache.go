package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.entries[key].val == nil {
		return nil, false
	}
	value := c.entries[key].val
	return value, true
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries:  map[string]cacheEntry{},
		mutex:    sync.Mutex{},
		interval: interval,
	}
	go newCache.reapLoop()
	return &newCache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, cacheEntry := range c.entries {
			if c.interval < time.Since(cacheEntry.createdAt) {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}

}
