package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Person JSONデコード用に構造体定義
type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func main() {
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
		fmt.Println(p)
		fmt.Printf("%d : %s :%s\n", p.ID, p.Name, p.Birthday)
	}
}
