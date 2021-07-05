package main

import (
	"fmt"
	"time"
)

func printHello(i int) {
	for {
		fmt.Printf("Hello from gorountine %d\n", i)
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		go printHello(i) //如果没有print就会一直被霸占
	}

	time.Sleep(time.Millisecond)
}
