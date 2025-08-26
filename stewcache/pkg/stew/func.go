package stew

import "time"

// has_data_expired checks if a given expiry time has passed.
//
// Parameters:
//   - expiry_time: The time at which the data should be considered expired
//
// Returns:
//   - bool: true if the current time is equal to or after the expiry_time, false otherwise
//
// This function is used internally by the cache to determine if an entry has expired.
// It's a utility function that simplifies the time comparison logic.
func has_data_expired(expiry_time time.Time) bool {
	current_time := time.Now()
	// If expiry_time is before or equal to current_time, the data has expired
	return !expiry_time.After(current_time)
}
