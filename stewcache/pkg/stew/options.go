package stew

import (
	"sync"
	"time"
)

type Config struct {
	TTL            time.Duration
	Interval       time.Duration
	BackupData     bool
	BackupInterval time.Duration
	BackupPath     string
}

// CacheValue represents a value stored in the cache along with its expiration time.
// It is a generic container that can hold any Go value.
type CacheValue struct {
	// Value is the actual data being cached. It can be of any type.
	Value any

	// TTL (Time To Live) specifies when this cache entry should expire.
	// Users must specify 0 value else itll pick the default TTL
	TTL time.Time
}

// Stew is the main cache structure that holds all cache data and synchronization primitives.
// It provides thread-safe operations for getting, setting, and deleting cache entries.
type Stew struct {
	// Mu is a read-write mutex that protects access to the GlobalCache.
	// It allows multiple concurrent readers or a single writer.
	Mu sync.RWMutex

	// GlobalCache is the underlying map that stores all cache entries.
	// It maps string keys to CacheValue structures.
	GlobalCache GlobalCache

	// stopCH is a channel used to signal the background cleanup goroutine to stop.
	stopCH chan struct{}

	// Settings for when the cache is initialized
	Config Config
	// wg is used to wait for all background goroutines to complete during shutdown.
	wg sync.WaitGroup
}

// SetOptions contains the parameters needed to set a value in the cache.
type SetOptions struct {
	// Key is the unique identifier for the cache entry.
	// It cannot be empty.
	Key string

	// Data contains the value to be cached and its optional TTL.
	Data CacheValue
}

// GlobalCache is a type alias for a map that stores cache entries.
// It maps string keys to CacheValue structures and is the underlying data structure
// used by the Stew cache.
type GlobalCache map[string]CacheValue
