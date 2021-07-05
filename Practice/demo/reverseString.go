package main

import (
	"fmt"
)

func main() {
	str := []byte{'A', 'B', 'D', 'E', 'c', '2'}

	sLen := len(str)
	half := sLen / 2

	for i := range str {
		if i < half {
			str[i], str[sLen-i-1] = str[sLen-i-1], str[i]
		}
	}
	fmt.Printf("%s\n", str)
}
