package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "name", 1)

	var ch = make(chan []byte)
	go request(ctx, "https://www.baidu.com", ch)
	go request(ctx, "https://www.taobao.com", ch)
	go request(ctx, "https://www.google.com", ch)

	i := 0
	for {
		resp := <-ch
		fmt.Printf("%d\n", len(resp))
		// 接收3个之后关闭chan
		i++
		if i == 3 {
			close(ch)
			break
		}
	}
	// more code...
}

func request(ctx context.Context, apiUrl string, res chan []byte) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("timeout...")
			res <- nil
		}
	}()
	resp, err := http.Get(apiUrl)
	if err != nil {
		res <- nil
		return
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res <- nil
		return
	}
	// 模拟超时
	time.Sleep(time.Second * 10)
	res <- all
}
