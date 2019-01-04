package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (user *User) SayHello() {
	fmt.Println(user.Name, ": Hello.")
}

func main() {
	fmt.Println("hogehoge")
	yamada := User{Name: "Yamada", Age: 30}

	fmt.Println(yamada.Name)
	fmt.Println(yamada.Age)

	yamada.SayHello()
}
