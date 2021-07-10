package main

import (
	"fmt"
	"time"
)

func main() {
	parse, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-08-01 14:14:20", time.Local)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", parse)

	fmt.Printf("%v\n", parse.Truncate(1*time.Hour))

	timeAfter()
}

func timeAfter() {
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		<-c
	}()

	select {
	case c <- 1:
		fmt.Println("channel...")
	case <-time.After(5 * time.Second):
		close(c)
		fmt.Println("timeout...")
	}
}
