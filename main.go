package main

import (
	"sync"
	"time"
)

type MappingType struct {
	val map[int]int
	lock *sync.Mutex
}

func main() {
	mapping := &MappingType{
		val: make(map[int]int, 10),
		lock: &sync.Mutex{},
	}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				// mapping.lock.Lock()
				mapping.val[i] = i
				// mapping.lock.Unlock()
			}
		}()
	}

	time.Sleep(10e9)
}