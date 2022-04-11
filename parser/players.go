package parser

import (
	"encoding/json"
	"go-crawler/common"
	"strconv"
)

const (
	Score   = "得分"
	Rebound = "篮板"
	Steal   = "抢断"
	Block   = "盖帽"
	Assist  = "助攻"
)

func parsePlayers(content []byte, context common.Context) common.ParseResult {
	response := common.JsonStatsResponse{}
	err := json.Unmarshal(content, &response)
	common.PanicErr(err)

	stats := common.Stats{
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

	return common.ParseResult{
		Result: stats,
	}
}

func setValue(v *common.JsonStatsItem) common.StatsValue {
	value, err := strconv.ParseFloat(v.Value, 64)
	common.PanicErr(err)

	avg, err := strconv.ParseFloat(v.LeagueAvg, 64)
	common.PanicErr(err)

	max, err := strconv.ParseFloat(v.LeagueMax, 64)
	common.PanicErr(err)

	serial, err := strconv.Atoi(v.Serial)
	if err != nil {
		serial = -1
	}

	return common.StatsValue{
		Value:     value,
		LeagueAvg: avg,
		LeagueMax: max,
		Serial:    serial,
	}
}
