package parser

import (
	"imooc/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	file, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(file, "安静的雪")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age:        30,
		Height:     163,
		Weight:     55,
		Income:     "8001-12000元",
		Gender:     "女",
		Name:       "安静的雪",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "销售",
		Location:   "山东潍坊",
		Star:       "魔羯座",
		House:      "租房",
		Car:        "未购车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
