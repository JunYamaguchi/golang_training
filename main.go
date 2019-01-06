// should use "go run main.go user.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	//yamada := user{Name: "Yamada", Age: 30}
	//fmt.Println(yamada.Name)
	//fmt.Println(yamada.Age)
	//yamada.SayHello()

	//mockDB := map[int]string{}
	mockDB := make(map[int]string)
	// map = make(map[key]value)

	mockDB[1] = "hoge"
	fmt.Println(mockDB)
}
