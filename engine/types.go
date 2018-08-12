package engine

/**
定义所需要的返回结构体
 */
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requsets []Request
	Iterms   []interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
