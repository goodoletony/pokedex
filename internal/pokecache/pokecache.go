package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	cache   map[string]cacheEntry
	cacheMutex sync.Mutex
	interval   time.Duration
}

func NewCache(duration time.Duration) *Cache {
	c := &Cache{
	cache: make(map[string]cacheEntry),
	interval: duration,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheMutex.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.cache[key] = entry
	c.cacheMutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.cacheMutex.Lock()
	entry, ok := c.cache[key]
	c.cacheMutex.Unlock()
	return entry.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	for key := range c.cache {
		elapsed := time.Since(c.cache[key].createdAt)
		if elapsed > c.interval {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C{
		c.cacheMutex.Lock()
		c.reap(c.interval)
		c.cacheMutex.Unlock()
	}
}
