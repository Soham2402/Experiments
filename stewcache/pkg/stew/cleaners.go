package stew

import (
	"time"
)

// dishwasher is a background worker that periodically cleans up expired cache entries.
// It runs in a separate goroutine and is responsible for:
//   - Running the cleanup at regular intervals specified by the interval parameter
//   - Performing a final cleanup when the cache is being closed
//   - Properly cleaning up resources when stopped
//
// The function runs until the stopCH channel is closed it exits.
func (s *Stew) dishwasher(interval time.Duration) {
	defer s.wg.Done()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.clean()
		case <-s.stopCH:
			return
		}
	}
}

// clean removes all expired entries from the cache.
// It acquires a write lock on the cache to ensure thread safety during cleanup.
// This function is called periodically by the dishwasher goroutine and during cache shutdown.
func (s *Stew) clean() {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	for k, v := range s.GlobalCache {
		if has_data_expired(v.TTL) {
			delete(s.GlobalCache, k)
		}
	}
}



func (s *Stew) RestartDishwasher(newInterval time.Duration) {
	// signal the old dishwasher to stop
	close(s.stopCH)
	s.wg.Wait() // wait for old goroutine to exit

	// create a new stop channel
	s.stopCH = make(chan struct{})

	// update config
	s.Config.Interval = newInterval

	// spin up new goroutine
	s.wg.Add(1)
	go s.dishwasher(newInterval)
}