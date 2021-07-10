package main

import "time"

// 协程顺序打印a、b、c各100次
func main() {
	a, b, c := make(chan int), make(chan int), make(chan int)
	count := 0
	go func(i int) {
		for {
			if i >= 300 {
				break
			}
			<-a
			print("a、")
			b <- 1
			time.Sleep(time.Millisecond * 100)
			i++
		}
	}(count)
	go func(i int) {
		for {
			if i >= 300 {
				break
			}
			<-b
			print("b、")
			c <- 1
			time.Sleep(time.Millisecond * 100)
		}
	}(count)
	go func(i int) {
		for {
			if i >= 300 {
				break
			}
			<-c
			print("c、")
			a <- 1
			time.Sleep(time.Millisecond * 100)
		}
	}(count)
	a <- 1
	for {
	}
}
