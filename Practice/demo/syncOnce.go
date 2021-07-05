package main

import "sync"

var once sync.Once

func main() {
	doOnce()
	doOnce()
	doOnce()
}

func doOnce() {
	once.Do(func() {
		println("one")
	})
}
