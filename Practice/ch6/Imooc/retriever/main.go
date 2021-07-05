package main

import (
	"fmt"
	"imooc/retriever/mock"
	"imooc/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "This is imooc!"}
	fmt.Println(r)

	r = real.Retriever{}

	fmt.Println(download(r))
}
