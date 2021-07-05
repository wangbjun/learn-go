package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

/**
获取url内容
*/
func main() {

	start := time.Now()

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}

		//使用io writter
		//io.Copy(os.Stdout, resp.Body)

		fmt.Println(url + " => " + resp.Status)

		//b, err := ioutil.ReadAll(resp.Body)
		//
		//resp.Body.Close()
		//
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//}
		//
		//fmt.Printf("%s", b)
	}

	usage := time.Since(start).Seconds()

	fmt.Printf("%.2fs\n", usage)
}
