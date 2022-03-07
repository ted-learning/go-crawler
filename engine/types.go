package engine

type Request struct {
	Url        string
	Context    Context
	ParserFunc func([]byte, Context) ParseResult
}

type ParseResult struct {
	Requests []Request
	Result   interface{}
}

type Context map[string]interface{}
