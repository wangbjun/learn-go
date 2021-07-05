package parser

import (
	"imooc/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="http://album.zhenai.com/u/(\d+)" target="_blank">([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	allProfile := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range allProfile {
		name := string(m[2])
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: "http://album.zhenai.com/u/" + string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	allCity := cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range allCity {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
