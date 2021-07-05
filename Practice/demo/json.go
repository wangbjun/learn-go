package main

import (
	"fmt"
	"reflect"
)

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (r Item) Say() string {
	fmt.Printf("name = %s\n", r.Name)

	return r.Name
}

func main() {
	i := Item{
		Name:  "json",
		Price: 10,
	}
	v := reflect.ValueOf(i)

	// 第1个成员值
	fmt.Printf("%v\n", v.Field(0))

	// 第2个成员值
	fmt.Printf("%v\n", v.FieldByName("Price"))

	// 调用其方法
	fmt.Printf("%v\n", v.Method(0).Call(nil))

	//修改，注意这里传的是地址
	m := reflect.ValueOf(&i)

	m.Elem().Field(0).SetString("JSON")
	m.Elem().FieldByName("Price").SetInt(100)

	fmt.Printf("%v\n", i.Name)
	fmt.Printf("%v\n", i.Price)
}

func Add(a, b interface{}) interface{} {
	switch reflect.TypeOf(a).String() {
	case "int":
		return reflect.ValueOf(a).Int() + reflect.ValueOf(b).Int()
	case "float64":
		return reflect.ValueOf(a).Float() + reflect.ValueOf(b).Float()
	}
	return nil
}

func Tag() {

}
