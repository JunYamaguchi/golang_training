// Shoud use go run main.go user.go

package main

import (
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func (user *user) SayHello() {
	fmt.Println(user.Name, ": Hello.")
}
