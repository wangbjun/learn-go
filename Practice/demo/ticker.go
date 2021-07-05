package main

import (
	"log"
	"time"
)

func main() {
	t1 := time.NewTicker(10 * time.Second)
	t2 := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-t1.C:
			log.Println("执行10s")
			time.Sleep(25 * time.Second)
		case <-t2.C:
			log.Println("执行20s")
		}
	}
}
