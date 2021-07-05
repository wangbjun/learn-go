package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	for i := 1000; i < 1500; i++ {
		go request(i)
	}

	time.Sleep(time.Second * 5)
}

func request(i int) {
	url := "http://t.wibliss.com/" + strconv.Itoa(i)
	resp, err := http.Post("http://t.wibliss.com/api/v1/create",
		"application/x-www-form-urlencoded", strings.NewReader("url="+url))
	if err != nil {
		panic(err)
	}
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", res)
	resp.Body.Close()
}
