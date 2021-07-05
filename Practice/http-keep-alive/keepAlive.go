package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{}

	for {
		request, err := http.NewRequest("GET", "https://www.jianshu.com/p/49551bda6619", nil)
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Set("Connection", "keep-alive")

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(response.StatusCode)

		log.Printf("%v\n", response.Header)
		log.Printf("%v\n", request.Header)
		response.Body.Close()
		time.Sleep(time.Second)
	}

}
