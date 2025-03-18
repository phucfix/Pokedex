package pokecache

import (
    "sync"
    "time"
)

type Cache struct {
    mu    *sync.Mutex
    cache map[string]cacheEntry
}

// TODO: expose a NewCache() function that creates a new cache with a configurable interval (time.Duration).
func NewCache(interval time.Duration) Cache {
    c := Cache{
        mu: &sync.Mutex{},
        cache: make(map[string]cacheEntry),
    }
    go c.reapLoop(interval)

    return c
}

func (c *Cache) Add(key string, value []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = cacheEntry{
        createdAt: time.Now().UTC(),
        val: value,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    val, ok := c.cache[key]
    return val.val, ok
}

// called when the cache is created
// Each time an interval (the time.Duration passed to NewCache) passes it should remove any entries that are older than the interval
//  This makes sure that the cache doesn't grow too large over time.
// For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.
func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

// 
func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    for k, v := range c.cache {
        if v.createdAt.Before(now.Add(-last)) {
            delete(c.cache, k)
        }
    }
}
