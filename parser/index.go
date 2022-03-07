package parser

import (
	"fmt"
	"go-crawler/engine"
	"regexp"
)

const teamAjaxRequestRegex = "(n\\.ajax\\({url:\"//)([a-zA-Z0-9&=_/.?]+)(\")"

func ParseIndex(content []byte, _ engine.Context) engine.ParseResult {
	compiled := regexp.MustCompile(teamAjaxRequestRegex)
	submatch := compiled.FindSubmatch(content)

	request := engine.Request{
		Url:        fmt.Sprintf("https://%s", string(submatch[2])),
		ParserFunc: parseTeams,
	}
	return engine.ParseResult{
		Requests: []engine.Request{request},
		Result:   fmt.Sprintf("Find teams ajax request url: %s", request.Url),
	}
}
