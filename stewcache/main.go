// Package main provides a simple example of using the Stew cache.
//
// This example demonstrates the basic usage of the Stew cache, including:
// - Creating a new cache instance
// - Setting a value with a key
// - Retrieving a value by key
// - Deleting a value by key
package main

import (
	"fmt"
	"time"

	"github.com/soham2402/stewcache/pkg/stew"
)

// main is the entry point of the application.
// It demonstrates the basic usage of the Stew cache.
func main() {
	// Create a new instance of the Stew cache
	newStew := stew.CreateStew()
	config := stew.NewConfig(
		stew.WithBackupData(true),
		stew.WithBackupInterval(1*time.Minute),
	)
	newStew.Config = *config
	// Ensure resources are cleaned up when we're done
	defer newStew.Close()

	// Create a cache entry with a key "test" and value 12323
	opt := stew.SetOptions{
		Key: "test",
		Data: stew.CacheValue{
			Value: 12323,
			// TTL is not set, so it will use the default (Jan 1, 2026)
		},
	}

	// Add the entry to the cache
	_, err := newStew.Set(&opt)
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}

	// Retrieve the value from the cache
	data, exists, err := newStew.Get("test")
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}
	fmt.Printf("Value: %v, Exists: %v\n", data, exists)

	// Delete the value from the cache
	_, err = newStew.Delete("test")
	if err != nil {
		fmt.Println("Error deleting value:", err)
		return
	}

	// Try to retrieve the deleted value
	data, exists, err = newStew.Get("test")
	if err != nil {
		fmt.Println("Error getting value after deletion:", err)
		return
	}
	fmt.Printf("After deletion - Value: %v, Exists: %v\n", data, exists)
}
