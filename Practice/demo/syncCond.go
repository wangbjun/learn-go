package main

import (
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})

func main() {
	go func() {
		println("1")
	}()
	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			cond.Wait()
			println(i)
			cond.L.Unlock()
		}(i)
	}

	// 确保所有协程启动完毕
	time.Sleep(time.Second * 1)

	cond.Broadcast()

	// 确保结果有时间输出
	time.Sleep(time.Second * 1)
}
