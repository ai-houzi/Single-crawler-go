package parser

import (
	"myProject/Single-crawler-go/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParserCity(contents []byte) engine.ParseResult {
	//fmt.Printf("%s\n",contents)

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		result.Iterms = append(result.Iterms, "user:"+name)
		result.Requsets = append(result.Requsets,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParserProfile(bytes, name)
				},
			})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requsets = append(result.Requsets,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParserCity,
			})
	}

	return result
}
