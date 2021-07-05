package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	t := time.NewTimer(5 * time.Second)
	select {
	case <-ch1:
		t.Stop()
	case <-t.C:
		fmt.Println("5s已经过去，任务超时")
	}
}
