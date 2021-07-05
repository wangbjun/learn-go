package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World!")

	var s, sep string

	//for i:=1;i<len(os.Args);i++ {
	//	s += sep+os.Args[i]
	//	sep = " "
	//}

	for i, args := range os.Args[1:] {
		println(i)
		s += sep + args
		sep = " "
	}

	fmt.Println(os.Args[1:])

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(s)
}
