package stew

import (
	"time"
)

// Set adds or updates a key-value pair in the cache with optional TTL.
//
// Parameters:
//   - opt: A pointer to SetOptions containing the key and data to be stored
//
// Returns:
//   - bool: true if the key was created, false if it was updated
//   - error: non-nil if an error occurred during the operation
//
// The function performs the following validations:
//   - Ensures the key is not empty
//   - Validates the TTL (if provided) is not in the past
//   - Sets a default TTL based on the config
//
// The operation is thread-safe and uses a write lock to prevent data races.
func (s *Stew) Set(opt *SetOptions) (bool, error) {
	cache_key := opt.Key
	if cache_key == "" {
		err := ErrEmptyKey
		return false, err
	}

	TTL := opt.Data.TTL
	expired := !TTL.IsZero() && has_data_expired(TTL)

	if expired {
		err := ErrExpiredTTL
		return false, err
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()

	// Set default TTL if none provided
	if TTL.IsZero() {
		opt.Data.TTL = time.Now().Add(s.Config.TTL)
	}

	_, exists := s.GlobalCache[cache_key]
	s.GlobalCache[cache_key] = opt.Data
	return !exists, nil
}

// Delete removes a key-value pair from the cache.
//
// Parameters:
//   - Key: The key of the item to delete
//
// Returns:
//   - bool: Always returns true if no error occurred
//   - error: non-nil if the key is empty
//
// The operation is thread-safe and uses a write lock.
func (s *Stew) Delete(Key string) (bool, error) {
	if Key == "" {
		err := ErrEmptyKey
		return false, err
	}

	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.GlobalCache, Key)
	return true, nil
}

// Get retrieves a value from the cache by its key.
//
// Parameters:
//   - Key: The key of the item to retrieve
//
// Returns:
//   - any: The stored value if found, nil otherwise
//   - bool: true if the key exists and is not expired, false otherwise
//   - error: non-nil if the key is empty
//
// The function performs the following checks:
//   - Validates the key is not empty
//   - Checks if the item has expired (in which case it's treated as non-existent)
//
// The operation is thread-safe and uses a read lock for concurrent access.
func (s *Stew) Get(Key string) (any, bool, error) {
	if Key == "" {
		err := ErrEmptyKey
		return nil, false, err
	}

	s.Mu.RLock()
	defer s.Mu.RUnlock()

	Data, exists := s.GlobalCache[Key]
	if exists {
		if has_data_expired(Data.TTL) {
			return nil, false, nil
		}
		return Data.Value, true, nil
	}

	return nil, false, nil
}
