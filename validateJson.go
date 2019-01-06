package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/go-playground/validator.v9"
)

// Person JSONデコード用に構造体定義
type Person struct {
	ID       *int    `json:"id" validate:"required"`
	Name     *string `json:"name" validate:"required"`
	Birthday string  `json:"birthday" validate:"required"`
	//Hoge     string `validate:"required"`
}

// 値がないときはゼロ値が入る
// 嫌なときは構造体のフィールドの型をポインタにして処理もポインタにするとnilが入る
// あるいはゼロ値なら無視をするアノテーションもあるomitempty?

func main() {
	v := validator.New()

	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("vro.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	for _, p := range persons {
		//fmt.Println(p)
		fmt.Printf("%d : %s :%s\n", *p.ID, *p.Name, p.Birthday)
		if err := v.Struct(p); err == nil {
			fmt.Println(*p.Name, "Valid")
		} else {
			fmt.Println(*p.Name, "Invalid:", err)
		}
	}

	fmt.Println("TypeOf:", reflect.TypeOf(persons))

	// []main.Person と型が返る
	if reflect.TypeOf(persons) == reflect.TypeOf([]Person{}) {
		fmt.Println("reflect check is true")
	}
}
