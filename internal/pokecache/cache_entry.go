package pokecache

import (
    "time"
)

type cacheEntry struct {
    createdAt time.Time
    val       []byte     // represents the raw data we're caching.
}

