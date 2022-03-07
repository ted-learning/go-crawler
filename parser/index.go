package parser

import (
	"fmt"
	"go-crawler/engine"
	"regexp"
)

const teamsLinkRegex = "(https://[a-zA-Z0-9&=_/.?]+)(['+s.teamId]*)([&cid=0-9]*)"
const teamAjaxRequestRegex = "(n\\.ajax\\({url:\"//)([a-zA-Z0-9&=_/.?]+)(\")"
const linkFormatKey = "linkFormat"

func ParseIndex(content []byte, _ engine.Context) engine.ParseResult {
	compiled := regexp.MustCompile(teamsLinkRegex)
	s := compiled.FindSubmatch(content)
	linkFormat := string(s[1]) + "%s" + string(s[3])
	context := map[string]interface{}{
		linkFormatKey: linkFormat,
	}

	compiled = regexp.MustCompile(teamAjaxRequestRegex)
	submatch := compiled.FindSubmatch(content)

	request := engine.Request{
		Url:        fmt.Sprintf("https://%s", string(submatch[2])),
		ParserFunc: parseTeams,
		Context:    context,
	}
	return engine.ParseResult{
		Requests: []engine.Request{request},
		Result:   fmt.Sprintf("Find teams ajax request url: %s", request.Url),
	}
}
