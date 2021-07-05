package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" //UTF-8
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))

	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println()

	fmt.Printf("Rune count: %d\n", utf8.RuneCount([]byte(s)))

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}

}
