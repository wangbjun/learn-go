package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerC(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*
		time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerC(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
