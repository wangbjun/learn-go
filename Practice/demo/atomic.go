package main

import (
	"sync"
	"sync/atomic"
)

type Person struct {
	num int64
}

var p = Person{num: 0}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go workerAtomic(&wg)
	go workerAtomic(&wg)
	go workerAtomic2(&wg)
	go workerAtomic2(&wg)
	wg.Wait()

	println(p.num)
}

func workerAtomic(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 1000020; i++ {
		atomic.AddInt64(&p.num, 1)
	}
}

func workerAtomic2(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt64(&p.num, -1)
	}
}
