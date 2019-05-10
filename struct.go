package main

import (
	"fmt"
)

const (
	A = 1
)

var b = 1

type User struct {
	name string
}

type Printer interface{
	sayName() string
}

func init() {
	fmt.Println("init")
}

func main() {
	user := &User {
		name: "jack",
	}
	print(user)
}

func (user *User) sayName() string {
	return user.name
}

func print(printer Printer) {
	fmt.Println(printer.sayName())
}
