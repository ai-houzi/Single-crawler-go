package main

import (
	"myProject/Single-crawler-go/engine"
	"myProject/Single-crawler-go/scheduler"
	"myProject/Single-crawler-go/zhenai/parser"
	"myProject/Single-crawler-go/persisit"
)

func main() {

	//多线程版
	city()

	//单线程版
	//simple()

}

func simple() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

func concurrentQueue() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 20,
		ItemChan:persisit.ItermSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

func concurrentSimple() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 20,
		ItemChan:persisit.ItermSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

func city()  {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 20,
		ItemChan:persisit.ItermSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParserCity,
	})
}