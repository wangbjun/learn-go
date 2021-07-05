package main

import (
	"fmt"
	"regexp"
)

const text = "My email is wangbenjun@gmail.com, and his email is benjun@baidu.com"

func main() {
	compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+).([a-zA-Z0-9]+)`)

	str := compile.FindAllStringSubmatch(text, -1)

	for _, j := range str {
		fmt.Println(j)
	}
}
