package engine

import (

	"log"
	"myProject/Single-crawler-go/fetcher"
)

func Run(sees ...Request) {
	var requests []Request

	for _, r := range sees {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching: %s", r.Url)
		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("fetcher error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		requests = append(requests, parseResult.Requsets...)

		for _, item := range parseResult.Iterms {
			log.Printf("got item %v", item)
		}
	}

}
