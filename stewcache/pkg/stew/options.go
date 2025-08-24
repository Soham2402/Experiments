package stew

import (
	"sync"
	"time"
)

type CacheValue struct {
	Value any
	TTL time.Time
}

type Stew struct {
	/// entry point of every single thing
	// define the settings needed to be initialized
	/// Global cache will be the global level map where every single Key Value pair will be used
	Mu sync.RWMutex
	GlobalCache GlobalCache
}

type SetOptions struct {
	Key string
	Data CacheValue
}


type GlobalCache map[string]CacheValue 