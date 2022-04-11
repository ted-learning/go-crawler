package parser

import (
	common2 "crawler-engine/common"
	"encoding/json"
	"strconv"
)

const (
	Score   = "得分"
	Rebound = "篮板"
	Steal   = "抢断"
	Block   = "盖帽"
	Assist  = "助攻"
)

func parsePlayers(content []byte, context common2.Context) common2.ParseResult {
	response := common2.JsonStatsResponse{}
	err := json.Unmarshal(content, &response)
	common2.PanicErr(err)

	stats := common2.Stats{
		PlayerId: context["playerId"].(string),
	}

	for _, v := range response.Data.StatsCompare {
		switch v.Type {
		case Score:
			stats.Score = setValue(&v)
		case Rebound:
			stats.Rebound = setValue(&v)
		case Steal:
			stats.Steal = setValue(&v)
		case Block:
			stats.Block = setValue(&v)
		case Assist:
			stats.Assist = setValue(&v)
		}
	}

	return common2.ParseResult{
		Result: stats,
	}
}

func setValue(v *common2.JsonStatsItem) common2.StatsValue {
	value, err := strconv.ParseFloat(v.Value, 64)
	common2.PanicErr(err)

	avg, err := strconv.ParseFloat(v.LeagueAvg, 64)
	common2.PanicErr(err)

	max, err := strconv.ParseFloat(v.LeagueMax, 64)
	common2.PanicErr(err)

	serial, err := strconv.Atoi(v.Serial)
	if err != nil {
		serial = -1
	}

	return common2.StatsValue{
		Value:     value,
		LeagueAvg: avg,
		LeagueMax: max,
		Serial:    serial,
	}
}
