package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/"))
	})
	http.HandleFunc("/hello/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/hello/"))
	})
	http.HandleFunc("/hello/abc/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/hello/abc/"))
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
