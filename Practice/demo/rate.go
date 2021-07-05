package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	for {
		if limiter.Allow() {
			fmt.Printf("%s: allow\n", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
