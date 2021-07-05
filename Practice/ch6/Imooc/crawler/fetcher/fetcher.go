package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	client := http.Client{}
	request, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	request.Header.Set("Cookie", "sid=NFVMYcVpYKKKgfsHxmBs")

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Got Wrong Status: %d", resp.StatusCode))
	}

	newReader := bufio.NewReader(resp.Body)
	bytes, err := newReader.Peek(1024)
	if err != nil {
		panic(err)
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	reader := transform.NewReader(resp.Body, encoding.NewDecoder())
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return all, nil
}
