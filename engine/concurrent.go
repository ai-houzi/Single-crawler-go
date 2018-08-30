package engine

import (
	"log"
	"myProject/Single-crawler-go/model"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifiter
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifiter interface {
	WorkReady(chan Request)
}

/**
多线程版本
 */
func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("duplicate url : %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Iterms {
			if _, ok := item.(model.Profile);ok{
				go func() {
					e.ItemChan <- item
				}()
			}

		}

		for _, request := range result.Requsets {
			if isDuplicate(request.Url) {
				log.Printf("duplicate url : %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, r ReadyNotifiter) {
	go func() {
		for {
			r.WorkReady(in)
			request := <-in
			result, err := work(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false

}
