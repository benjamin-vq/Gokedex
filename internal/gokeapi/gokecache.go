package gokeapi

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry
	rw    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) *Cache {

	cache := Cache{
		Cache: map[string]cacheEntry{},
		rw:    &sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return &cache

}

func (c *Cache) Add(key string, value []byte) {
	log.Printf("Adding the following key to cache: '%s'\n", key)

	entry := cacheEntry{
		time.Now(),
		value,
	}

	c.rw.Lock()
	defer c.rw.Unlock()

	c.Cache[key] = entry
}

func (c *Cache) Get(key string) (value []byte, found bool) {
	log.Printf("Getting the following key from cache: '%s'\n", key)

	c.rw.RLock()
	defer c.rw.RUnlock()

	entry, present := c.Cache[key]

	if !present {
		return nil, false
	}

	return entry.value, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	// This sucks. Since we are checking every 'interval', an entry
	// might be created just after the ticker channel receives a tick,
	// which means the entry won't be stale until the next tick, so in theory
	// an entry's lifetime is (interval, 2*interval)
	for range ticker.C {
		c.rw.Lock()

		log.Println("Interval passed, checking cache for stale entries")

		now := time.Now()
		for key, entry := range c.Cache {

			if now.Sub(entry.createdAt)-interval.Abs() > 0 {
				log.Printf("Removing stale entry of key: %s (it was created %v seconds ago)\n",
					key, now.Sub(entry.createdAt))

				delete(c.Cache, key)
			} else {
				log.Printf("Entry of key %s was created %v seconds ago (not stale)",
					key, now.Sub(entry.createdAt))
			}
		}

		c.rw.Unlock()
	}
}
