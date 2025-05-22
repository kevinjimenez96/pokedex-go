package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(d time.Duration) Cache {
	var cache = Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}
	go cache.reapLoop(d)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	//fmt.Println("cache added")
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if entry, ok := c.cache[key]; ok {
		//fmt.Println("cache used")
		return entry.val, true
	} else {
		return nil, false
	}

}

func (c *Cache) reapLoop(d time.Duration) {
	ticker := time.NewTicker(d)
	for range ticker.C {
		c.reap(d)
	}
}

func (c *Cache) reap(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, val := range c.cache {
		if time.Since(val.createdAt) > d {
			delete(c.cache, key)
			//fmt.Println("cache deleted")
		}
	}
}
