package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestDivHandler(t *testing.T) {
	recorder := httptest.NewRecorder()

	params := url.Values{}
	params.Add("a", "42")
	params.Add("b", "2")

	request, _ := http.NewRequest("POST", "/div", strings.NewReader(params.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	DivHandler(recorder, request)

	if recorder.Result().StatusCode != 200 {
		t.Error("Test failed")
	}

	body, _ := ioutil.ReadAll(recorder.Result().Body)
	if string(body) != "21" {
		t.Error("Test failed")
	}
}
