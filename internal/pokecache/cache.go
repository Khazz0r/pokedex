package pokecache

import (
	"time"
	"sync"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val []byte
}

type Cache struct {
	sync.RWMutex
	Entries map[string]CacheEntry
	Timeout time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries: make(map[string]CacheEntry),
		Timeout: interval,
	}

	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()
	entry := CacheEntry{
		CreatedAt: time.Now(),
		Val: val,
	}
	c.Entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()
	if entry, exists := c.Entries[key]; exists {
		return entry.Val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Timeout)
	defer ticker.Stop()
	for range ticker.C {
		c.Lock()
		func() {
			defer c.Unlock()
			for key, entry := range c.Entries {
				if time.Since(entry.CreatedAt) > c.Timeout {
					delete(c.Entries, key)
				}
			}
		}()
	}
}
