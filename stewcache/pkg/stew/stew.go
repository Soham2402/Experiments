package stew

import (
	"time"
)

// CreateStew initializes and returns a new instance of Stew cache.
// It sets up the necessary data structures and starts the background cleanup goroutine.
// The function returns a pointer to the newly created Stew instance.
//
// Example:
//
//	cache := CreateStew()
//	defer cache.Close()
func CreateStew() *Stew {
	gc := make(GlobalCache)
	stopCH := make(chan struct{})
	new_stew := Stew{GlobalCache: gc, stopCH: stopCH}
	new_stew.wg.Add(1)
	go new_stew.dishwasher(1 * time.Second)
	return &new_stew
}

// Close gracefully shuts down the cache.
// It stops the background cleanup goroutine and waits for any pending operations to complete.
// This function should be called when the cache is no longer needed to prevent goroutine leaks.
//
// Example:
//
//	cache := CreateStew()
//	defer cache.Close()
func (s *Stew) Close() {
	close(s.stopCH)
	s.wg.Wait()
}
