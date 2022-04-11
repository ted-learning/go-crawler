package parser

import (
	common2 "crawler-engine/common"
	"encoding/json"
	"fmt"
)

const playerDetailsTemp = "https://matchweb.sports.qq.com/player/stats?&callback=playerStats&playerId=%s&from=web"

func parseRosters(content []byte, _ common2.Context) common2.ParseResult {
	rosters := common2.JsonRosters{}
	err := json.Unmarshal(content, &rosters)
	common2.PanicErr(err)

	if rosters.Code != 0 {
		common2.PanicErr(fmt.Errorf("parse roster error, code: %d", rosters.Code))
	}

	var requests []common2.Request
	for _, v := range rosters.Data {
		requests = append(requests, common2.Request{
			Url:        fmt.Sprintf(playerDetailsTemp, v.PlayerId),
			ParserFunc: parsePlayers,
			Context:    map[string]interface{}{"playerId": v.PlayerId},
		})
	}

	return common2.ParseResult{
		Requests: requests,
		Result:   rosters.Data,
	}
}
