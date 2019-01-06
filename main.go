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
	fmt.Println("mockDB is", mockDB)
	fmt.Println("mockDB[1] is ", mockDB[1])
	fmt.Println("mockDB[2] is ", mockDB[2])
	// 定義してない値を開こうとするとvalueのゼロ値が返る
	//fmt.Println("mockDB[\"hoge\"] is", mockDB["hoge"])
	// keyの方が合わないと開けないしコンパイルエラーが出る

	var (
		mockDBValue string
		ok          bool
	)

	mockDBValue, ok = mockDB[1]
	fmt.Println(mockDBValue, ok)

	if ok {
		fmt.Println("mockDB[1] exist check is True")
	} else {
		fmt.Println("mockDB[1] exits check is False")
	}

	mockDBValue, ok = mockDB[2]
	fmt.Println(mockDBValue, ok)

	if ok {
		fmt.Println("mockDB[2] exist check is True")
	} else {
		fmt.Println("mockDB[2] exits check is False")
	}

	delete(mockDB, 1)
	fmt.Println("mockDB[1] is deleted")

	mockDBValue, ok = mockDB[1]
	fmt.Println(mockDBValue, ok)

	if ok {
		fmt.Println("mockDB[1] exist check is True")
	} else {
		fmt.Println("mockDB[1] exits check is False")
	}
}
