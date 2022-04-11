package parser

import (
	common2 "crawler-engine/common"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

const (
	EAST      = 1
	EASTSOUTH = 11
	ATLANTIC  = 12
	CENTRAL   = 13
	WEST      = 2
	PACIFIC   = 21
	WESTSOUTH = 22
	WESTNORTH = 23
)

const teamLinkTemp = "https://matchweb.sports.qq.com/team/players?teamId=%s&competitionId=100000"

func parseTeams(content []byte, _ common2.Context) common2.ParseResult {
	jsonNBA := parseJson(content)
	emptyTeams := make([]common2.JsonTeam, 0)
	nba := common2.NBA{
		East: common2.East{
			EastSouth: convertToTeam(&jsonNBA.EastSouth, EAST, EASTSOUTH, &jsonNBA.East),
			Central:   convertToTeam(&jsonNBA.Central, EAST, CENTRAL, &jsonNBA.East),
			Atlantic:  convertToTeam(&jsonNBA.Atlantic, EAST, ATLANTIC, &jsonNBA.East),
			Total:     convertToTeam(&jsonNBA.East, EAST, 0, &emptyTeams),
		},

		West: common2.West{
			Pacific:   convertToTeam(&jsonNBA.Pacific, WEST, PACIFIC, &jsonNBA.West),
			WestNorth: convertToTeam(&jsonNBA.WestNorth, WEST, WESTNORTH, &jsonNBA.West),
			WestSouth: convertToTeam(&jsonNBA.WestSouth, WEST, WESTSOUTH, &jsonNBA.West),
			Total:     convertToTeam(&jsonNBA.West, WEST, 0, &emptyTeams),
		},
	}
	result := common2.ParseResult{Result: nba}

	for _, v := range nba.East.Total {
		v.Link = fmt.Sprintf(teamLinkTemp, v.TeamId)
		result.Requests = append(result.Requests, common2.Request{
			Url:        v.Link,
			ParserFunc: parseRosters,
		})
	}

	for _, v := range nba.West.Total {
		v.Link = fmt.Sprintf(teamLinkTemp, v.TeamId)
		result.Requests = append(result.Requests, common2.Request{
			Url:        v.Link,
			ParserFunc: parseRosters,
		})
	}

	return result
}

func findTemRank(team *common2.JsonTeam, total *[]common2.JsonTeam) int {
	for _, i := range *total {
		if i.TeamId == team.TeamId {
			rank, err := strconv.Atoi(i.Serial)
			common2.PanicErr(err)
			return rank
		}
	}
	rank, err := strconv.Atoi(team.Serial)
	common2.PanicErr(err)
	return rank
}

func findAreaName(id int) string {
	switch id {
	case EAST:
		return "东部"
	case EASTSOUTH:
		return "东南赛区"
	case ATLANTIC:
		return "大西洋赛区"
	case CENTRAL:
		return "中部赛区"
	case WEST:
		return "西部"
	case PACIFIC:
		return "太平洋赛区"
	case WESTSOUTH:
		return "西南赛区"
	case WESTNORTH:
		return "西北赛区"
	default:
		return ""
	}
}

func parseJson(content []byte) common2.JsonNBA {
	var tempMap []interface{}
	err := json.Unmarshal(content, &tempMap)
	common2.PanicErr(err)

	marshal, err := json.Marshal(tempMap[1])
	common2.PanicErr(err)

	jsonNBA := common2.JsonNBA{}
	err = json.Unmarshal(marshal, &jsonNBA)
	common2.PanicErr(err)
	return jsonNBA
}

func convertToTeam(source *[]common2.JsonTeam, areaId, divId int, total *[]common2.JsonTeam) []common2.Team {
	target := make([]common2.Team, len(*source))
	for i, v := range *source {
		wins, err := strconv.Atoi(v.Wins)
		common2.PanicErr(err)

		losses, err := strconv.Atoi(v.Losses)
		common2.PanicErr(err)

		winingPercentage, err := strconv.Atoi(v.WiningPercentage)
		common2.PanicErr(err)

		divRank, err := strconv.Atoi(v.DivRank)
		common2.PanicErr(err)

		rank := findTemRank(&v, total)

		target[i] = common2.Team{
			Badge:            v.Badge,
			Name:             v.Name,
			TeamId:           v.TeamId,
			Wins:             wins,
			Losses:           losses,
			WiningPercentage: winingPercentage,
			Rank:             rank,
			Area:             findAreaName(areaId),
			AreaId:           areaId,
			Div:              findAreaName(divId),
			DivId:            divId,
			DivRank:          divRank,
		}
	}

	sort.Slice(target, func(i, j int) bool {
		return target[i].Rank < target[j].Rank
	})

	return target
}
