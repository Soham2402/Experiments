package stew

import "time"


func has_data_expired(expiry_time time.Time) bool {
	current_time := time.Now()
	if expiry_time.After(current_time){
		return false
	}else{
		return true
	}
}

