package main

import (
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/div", DivHandler)
	_ = http.ListenAndServe(":8888", nil)
}

func DivHandler(writer http.ResponseWriter, request *http.Request) {
	a := request.PostFormValue("a")
	b := request.PostFormValue("b")

	paramA, _ := strconv.Atoi(a)
	paramB, _ := strconv.Atoi(b)

	i, err := Div(paramA, paramB)
	if err != nil {
		_, _ = writer.Write([]byte("error"))
		return
	}
	_, _ = writer.Write([]byte(strconv.Itoa(i)))
}
