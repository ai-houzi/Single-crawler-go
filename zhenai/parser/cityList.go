package parser

import (

	"regexp"
	"myProject/Single-crawler-go/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

/**
	正则表达式获取城市列表
 */
func ParseCityList(contents []byte) engine.ParseResult {
	r := regexp.MustCompile(cityListRe)

	matches := r.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {

		result.Iterms = append(result.Iterms, "city:"+string(m[2]))
		result.Requsets = append(result.Requsets,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParserCity,
			})
	}

	return result
}
