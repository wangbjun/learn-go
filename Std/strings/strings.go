package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	replace()
	fmt.Println(strings.Count("abc", ""))
	fmt.Println(strings.Count("abc", "b"))

	fmt.Println(strings.Fields(" bd abc 12"))
	fmt.Println(strings.FieldsFunc(" bd abc 35", unicode.IsSpace))

	fmt.Println(strings.Split(" bd abc 35", ""))

	fmt.Println(strings.HasPrefix("bd abc 35", "bd"))
}

func replace() {
	println(strings.Replace("abc abc abc", "ab", "99", -1))

	replacer := strings.NewReplacer("a", "a1", "b", "b1")
	s := replacer.Replace("abc abc abc")
	println(s)
}
