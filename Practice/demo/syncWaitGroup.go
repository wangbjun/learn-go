package main

import "sync"

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 计数+1
		go func() {
			println("1")
			wg.Done() // 计数-1，相当于wg.add(-1)
		}()
	}
	wg.Wait() // 阻塞带等待所有协程执行完毕
}
