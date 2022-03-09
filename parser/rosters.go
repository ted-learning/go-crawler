package parser

import (
	"encoding/json"
	"fmt"
	"go-crawler/common"
)

type JsonRosters struct {
	Code int      `json:"code"`
	Data []Player `json:"data"`
}

type Player struct {
	PlayerId  string `json:"playerId"`
	CnName    string `json:"cnName"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Logo      string `json:"logo"`
	Position  string `json:"position"`
	JerseyNum string `json:"jerseyNum"`
}

const playerDetailsTemp = "https://matchweb.sports.qq.com/player/stats?&callback=playerStats&playerId=%s&from=web"

func parseRosters(content []byte, _ common.Context) common.ParseResult {
	rosters := JsonRosters{}
	err := json.Unmarshal(content, &rosters)
	common.PanicErr(err)

	if rosters.Code != 0 {
		common.PanicErr(fmt.Errorf("parse roster error, code: %d", rosters.Code))
	}

	var requests []common.Request
	for _, v := range rosters.Data {
		requests = append(requests, common.Request{
			Url:        fmt.Sprintf(playerDetailsTemp, v.PlayerId),
			ParserFunc: parsePlayers,
			Context:    map[string]interface{}{"playerId": v.PlayerId},
		})
	}

	return common.ParseResult{
		Requests: requests,
		Result:   rosters.Data,
	}
}
