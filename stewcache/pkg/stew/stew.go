package stew

import (
	"fmt"
	"time"
)


func CreateStew() *Stew {
	gc := make(GlobalCache)
	new_stew := Stew{GlobalCache: gc}
	return &new_stew
}


// Sets or updates an existing Key
// returns false and error if an error
// returns true amd nil when created
// returns false and nil when updated 
func (s *Stew) Set(opt *SetOptions)(bool, error){
	cache_key := opt.Key
	var expired bool
	if cache_key == ""{
		err := fmt.Errorf("Failed to set Value, Key cannot be empty")
		return false, err
	}
	TTL := opt.Data.TTL
	is_zero := TTL.IsZero()
	if !is_zero {
		expired = has_data_expired(TTL)
	}else{
		expired = false
	}
	if expired{
		err := fmt.Errorf("Failed to set Value, date of expiry is invalid")
		return false, err
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if TTL.IsZero() {
		opt.Data.TTL = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	_, exists := s.GlobalCache[cache_key]
	s.GlobalCache[cache_key] = opt.Data
	return !exists, nil
}

func (s *Stew) Delete(Key string)(bool, error){
	if Key == ""{
		err := fmt.Errorf("Failed to delete Value, Key cannot be empty")
		return false, err
	}
	// Add a read lock so that multiple routines can read it parallely
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.GlobalCache, Key)
	return true, nil
}

// The Get operation allows you to get the Value wrt to a Key
// If the Key does not exist it will return nil.false
// If the Key exists with a nul Value it shall return nil, true
func (s *Stew) Get(Key string) (any, bool, error){
	var exists bool
	// Add a read lock so that multiple routines can read it parallely
	if Key == ""{
		err := fmt.Errorf("Failed to get Value, Key cannot be empty")
		return nil, false, err
	}
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	Data, exists := s.GlobalCache[Key]	
	if exists{
		expired := has_data_expired(Data.TTL)
		if expired{
			exists = false
			return nil, exists, nil
		}else{
			return Data.Value, exists, nil
		}
	}else{		
		return nil, false, nil
	}

}
