package stew

import (
	"fmt"
	"time"
)


func (s *Stew)compost(interval time.Duration) {
	defer s.wg.Done()
	ticker := time.NewTicker(interval)
	for {
		select {
		case <- ticker.C:
			s.collect()
		case <- s.stopCH:
			break
		}
		
	}
}


func (s *Stew) collect(){
	fmt.Println("Okay backup thingy is quiet cool")
}