package parser

import (
	"crawler-engine/common"
	"fmt"
	"regexp"
)

const teamAjaxRequestRegex = "(n\\.ajax\\({url:\"//)([a-zA-Z0-9&=_/.?]+)(\")"

func ParseIndex(content []byte, _ common.Context) common.ParseResult {
	compiled := regexp.MustCompile(teamAjaxRequestRegex)
	submatch := compiled.FindSubmatch(content)

	request := common.Request{
		Url:        fmt.Sprintf("https://%s", string(submatch[2])),
		ParserFunc: parseTeams,
	}
	return common.ParseResult{
		Requests: []common.Request{request},
		Result:   nil,
	}
}
