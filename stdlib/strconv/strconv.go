package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.ParseInt("250", 10, 32)
	if err != nil {
		panic(err)
	}
	println(i)

	println(strconv.FormatInt(255, 10))

	b, err := strconv.ParseBool("1")
	if err != nil {
		panic(err)
	}
	println(b)

	f, _ := strconv.ParseFloat("1.2344", 0)
	fmt.Printf("%f\n", f)

	float := strconv.FormatFloat(1.5445, 'f', 6, 64)

	println(float)

	fmt.Println("This is \"studygolang.com\" website")

	fmt.Println("This is", strconv.Quote("studygolang.com"), "website")
}
