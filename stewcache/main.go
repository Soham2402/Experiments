package main

import (
	"fmt"

	"github.com/soham2402/stewcache/pkg/stew"
)

func main(){	
	newStew := stew.CreateStew()
	opt := stew.SetOptions{Key: "test",Data: stew.CacheValue{Value:12323}}
	_,err := newStew.Set(&opt)
	if err != nil {
		fmt.Println(err)
	}
	// Data, exists, err := newStew.Get("test")
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(Data, exists)	
	newStew.Delete("test")
	Data, exists, err := newStew.Get("test")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(Data, exists)
}
