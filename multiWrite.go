package main

//import (
//	"sync"
//	"time"
//)
//
//func writeWithoutLock()  {
//	mapping := make(map[int]int, 0)
//	for i := 0; i < 4; i++ {
//		go func() {
//			for i := 0; i < 100;i++ {
//				mapping[i] = i
//			}
//		}()
//	}
//	time.Sleep(1e9)
//}
//
//type Mapping struct {
//	data map[int]int
//	lock *sync.Mutex
//}
//
//func writeWithLock()  {
//	mapping := &Mapping{
//		data: make(map[int]int),
//		lock: &sync.Mutex{},
//	}
//	for i := 0; i < 4; i++ {
//		go func() {
//			for i := 0; i < 100;i++ {
//				mapping.lock.Lock()
//				mapping.data[i] = i
//				mapping.lock.Unlock()
//			}
//		}()
//	}
//	time.Sleep(1e9)
//}
//
//func main() {
//	writeWithoutLock()
//	writeWithLock()
//}
