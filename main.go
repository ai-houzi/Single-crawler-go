package main

import (
	"myProject/Single-crawler-go/engine"
	"myProject/Single-crawler-go/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
