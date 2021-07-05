package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{}
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < 1000; i++ {
				go func(i int) {
					url := "http://yunqiblog.cn/"
					request, _ := http.NewRequest("GET", url, nil)
					request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
					do, err := client.Do(request)
					if err == nil {
						fmt.Printf("第 %d 次 --- %v\n", i, do.StatusCode)
					}
				}(i)
			}
			fmt.Println("Job Start...")
		}
	}
}
