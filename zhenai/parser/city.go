package parser

import (
	"myProject/Single-crawler-go/engine"
	"regexp"
)

const  cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) engine.ParseResult{
	//fmt.Printf("%s\n",contents)

	r := regexp.MustCompile(cityRe)

	matches := r.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {

		result.Iterms = append(result.Iterms, "user:"+string(m[2]))
		result.Requsets = append(result.Requsets,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParse,
			})
	}

	return result
}