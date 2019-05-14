package main

//import (
//	"fmt"
//	"time"
//)
//
//func serialCompute() {
//	sum := 0
//	start := time.Now()
//	for i := 0; i < 4e9; i++ {
//		sum += i
//	}
//	end := time.Now()
//	fmt.Println(end.Sub(start), sum)
//}
//
//func parallelCompute() {
//	count, sum, ch := 4, 0, make(chan int, 4)
//	start := time.Now()
//	for i := 0; i < count; i++ {
//		go func(k int) {
//			val := 0
//			for j := k * 1e9; j < (k+1)*1e9; j++ {
//				val += j
//			}
//			ch <- val
//		}(i)
//	}
//	for i := count; i > 0; i-- {
//		sum += <-ch
//	}
//	end := time.Now()
//	fmt.Println(end.Sub(start), sum)
//}
//
//func main() {
//	serialCompute()
//	parallelCompute()
//}
