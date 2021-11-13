package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"plugin"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		do()
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

func do() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %s", err)
		}
	}()
	a := gjson.ValidBytes([]byte("1111111"))
	println(a)
	open, err := plugin.Open("/home/jwang/Documents/Work/learn_go/test/plg/plg.so")
	if err != nil {
		panic(err)
	}
	lookup, err := open.Lookup("GetName")
	if err != nil {
		panic(err)
	}
	res := lookup.(func() string)()
	fmt.Printf("%v\n", res)

	name, err := open.Lookup("Name")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", *name.(*string))
}
