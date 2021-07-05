package main

import (
	"fmt"
	"sync"
)

type User struct {
	name string
}

var pool = sync.Pool{
	New: func() interface{} {
		return User{
			name: "default name",
		}
	},
}

func main() {
	pool.Put(User{name: "name1"})
	pool.Put(User{name: "name2"})

	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
	fmt.Printf("%v\n", pool.Get())
}
