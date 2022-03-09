package parser

import (
	"encoding/json"
	"go-crawler/common"
	"log"
	"strconv"
)

type JsonStatsResponse struct {
	Code int `json:"code"`
	Data struct {
		StatsCompare []JsonStatsItem `json:"statsCompare"`
	} `json:"data"`
}

type JsonStatsItem struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Serial    string `json:"serial"`
	LeagueAvg string `json:"leagueAvg"`
	LeagueMax string `json:"leagueMax"`
}

type Stats struct {
	PlayerId string
	Score    StatsValue
	Rebound  StatsValue
	Steal    StatsValue
	Block    StatsValue
	Assist   StatsValue
}

type StatsValue struct {
	Value     float64
	LeagueAvg float64
	LeagueMax float64
	Serial    int
}

const (
	Score   = "得分"
	Rebound = "篮板"
	Steal   = "抢断"
	Block   = "盖帽"
	Assist  = "助攻"
)

var playerStatsCount = 0

func parsePlayers(content []byte, context common.Context) common.ParseResult {
	response := JsonStatsResponse{}
	err := json.Unmarshal(content, &response)
	common.PanicErr(err)

	stats := Stats{
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
	playerStatsCount++
	log.Printf("==============player stats count :%d\n", playerStatsCount)

	return common.ParseResult{
		Result: stats,
	}
}

func setValue(v *JsonStatsItem) StatsValue {
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

	return StatsValue{
		Value:     value,
		LeagueAvg: avg,
		LeagueMax: max,
		Serial:    serial,
	}
}
