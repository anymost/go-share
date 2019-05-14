package main

import "fmt"

func main() {
	dist, src := make([]int, 4), []int{1, 2, 3, 4}
	copy(dist, src)
	fmt.Println(dist)
}
