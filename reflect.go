package main

//import "fmt"
//
//func add(a interface{}, b interface{}) (interface{}, bool) {
//	if intA, isOk := a.(int); isOk {
//		if intB, isOk := b.(int); isOk {
//			return intA + intB, true
//		} else {
//			return nil, false
//		}
//	}
//
//	if float64A, isOk := a.(float64); isOk {
//		if float64B, isOk := b.(float64); isOk {
//			return float64A + float64B, true
//		} else {
//			return nil, false
//		}
//	}
//
//	return nil, false
//}
//
//func main()  {
//	val, isOk := add(1, 2)
//	if isOk {
//		if intVal, isOk := val.(int); isOk {
//			fmt.Println(intVal)
//		}
//	}
//}