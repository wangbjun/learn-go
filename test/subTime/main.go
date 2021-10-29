package main

import (
	"fmt"
	"time"
)

func main() {
	parse, err := time.Parse("2006-01-02 15:04:05", "2021-10-26 12:00:00")
	if err != nil {
		panic(err)
	}
	sub := time.Now().Sub(parse)
	fmt.Printf("%v\n", sub < time.Minute*15)
}
