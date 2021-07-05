package engine

import (
	"imooc/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)

		if err != nil {
			continue
		}

		requests = append(requests, result.Requests...)
		if len(result.Items) > 0 {
			for _, v := range result.Items {
				log.Printf("Got items: %v\n", v)
			}
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fecth Url: %s\n", r.Url)

	content, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(content), nil
}
