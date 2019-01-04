// Shoud use go run main.go user.go

package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (user *User) SayHello() {
	fmt.Println(user.Name, ": Hello.")
}
