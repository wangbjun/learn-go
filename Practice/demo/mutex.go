package main

import "sync"

var lock = sync.Mutex{}

func main() {
	x := 0
	var wg sync.WaitGroup
	wg.Add(5)
	go worker(&x, &wg)
	go worker(&x, &wg)
	go worker(&x, &wg)
	go worker(&x, &wg)
	go worker(&x, &wg)
	wg.Wait()

	println(x)
}

func worker(x *int, group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		*x += 1
		lock.Unlock()
	}
}
