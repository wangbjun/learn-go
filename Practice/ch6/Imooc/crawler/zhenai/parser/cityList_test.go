package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCitys := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	requestCount := len(result.Requests)

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}

	if requestCount != resultSize {
		t.Errorf("result should have %d request, "+
			"but had %d", resultSize, requestCount)
	}
}
