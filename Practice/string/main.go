package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "我爱学习Go语言"

	rStr := []rune(str)

	fmt.Printf("%s\n", string(rStr[2:4])) //学习

	fmt.Printf("%s\n", string(rStr[3:])) //习Go语言

	fmt.Printf("%s\n", string(rStr[:3])) //我爱学

	r := "我爱学习Go语言"

	fmt.Printf("%s\n", r)
	fmt.Printf("%v\n", []byte(r))
	fmt.Printf("%v\n", []rune(r))

	subString := SubString(r, 1, 3)

	fmt.Printf("%s\n", subString)
	fmt.Printf("%s\n", ReverseString("1234567"))
	fmt.Printf("%s\n", ReverseRuneString(r))
}

func SubString(str string, start, end int) string {
	var n, i, k int

	for k = range str {
		if n == start {
			i = k
		}
		if n == end {
			break
		}
		n++
	}

	return str[i:k]
}

func ReverseString(str string) string {
	b := []byte(str)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func ReverseRuneString(s string) string {
	var start, size, end int
	buf := make([]byte, len(s))
	for end < len(s) {
		_, size = utf8.DecodeRuneInString(s[start:])
		end = start + size
		copy(buf[len(buf)-end:], s[start:end])
		start = end
	}
	return string(buf)
}
