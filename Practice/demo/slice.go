package main

import (
	"fmt"
	"net/http"
)

type myServer struct{}

func (myServer) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	_, _ = writer.Write([]byte("Hello"))
}

func main() {
	fmt.Println("before return")
	do()
	fmt.Println("after return")
	err := http.ListenAndServe(":8888", myServer{})
	if err != nil {
		panic(err)
	}
}

func do() {
	fmt.Println("do")
	return
}
