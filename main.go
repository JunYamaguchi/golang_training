// should use "go run main.go user.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hogehoge")
	yamada := User{Name: "Yamada", Age: 30}

	fmt.Println(yamada.Name)
	fmt.Println(yamada.Age)

	yamada.SayHello()
}
