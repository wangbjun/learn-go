package main

import (
	"fmt"
	"plugin"
)

func main() {
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
