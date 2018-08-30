package engine

import (
	"log"
	"myProject/Single-crawler-go/fetcher"
)

type SimpleEngine struct {}

/**
单线程版本
 */
func (e SimpleEngine) Run(sees ...Request) {
	var requests []Request

	for _, r := range sees {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := work(r)

		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requsets...)

		for _, item := range parseResult.Iterms {
			log.Printf("got item %v", item)
		}
	}

}

func work(r Request) (ParseResult, error) {
	log.Printf("fetching: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("fetcher error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
