package main

import (
	"encoding/json"
	"fmt"
)

type Goods struct {
	Name    string
	Price   int
	Address string `json:"address2"`
	Tag     string
}

func main() {
	jsonStr := `{"Name":"商品1","Price":100,"address2":"中国","Tag":"特价"}`

	goods := Goods{}

	err := json.Unmarshal([]byte(jsonStr), &goods)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", goods)
}
