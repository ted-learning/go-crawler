package parser

import (
	"encoding/json"
	"fmt"
	"go-crawler/common"
	"sort"
	"strconv"
)

type JsonNBA struct {
	//Rank
	East []JsonTeam `json:"east"`
	West []JsonTeam `json:"west"`
	//East
	EastSouth []JsonTeam `json:"eastsouth"`
	Atlantic  []JsonTeam `json:"atlantic"`
	Central   []JsonTeam `json:"central"`
	//West
	Pacific   []JsonTeam `json:"pacific"`
	WestSouth []JsonTeam `json:"westsouth"`
	WestNorth []JsonTeam `json:"westnorth"`
}

type JsonTeam struct {
	Badge            string `json:"badge"`
	Name             string `json:"name"`
	TeamId           string `json:"teamId"`
	Wins             string `json:"wins"`
	Losses           string `json:"losses"`
	WiningPercentage string `json:"wining-percentage"`
	DivRank          string `json:"divRank"`
	Serial           string `json:"serial"`
}

type NBA struct {
	East East
	West West
}

type East struct {
	EastSouth []Team
	Atlantic  []Team
	Central   []Team
	Total     []Team
}

type West struct {
	Pacific   []Team
	WestSouth []Team
	WestNorth []Team
	Total     []Team
}

type Team struct {
	Badge            string
	Name             string
	TeamId           string
	Wins             int
	Losses           int
	WiningPercentage int
	DivRank          int
	Rank             int
	Link             string
	Area             string
	AreaId           int
	Div              string
	DivId            int
}

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

func parseTeams(content []byte, _ common.Context) common.ParseResult {
	jsonNBA := parseJson(content)
	emptyTeams := make([]JsonTeam, 0)
	nba := NBA{
		East: East{
			EastSouth: convertToTeam(&jsonNBA.EastSouth, EAST, EASTSOUTH, &jsonNBA.East),
			Central:   convertToTeam(&jsonNBA.Central, EAST, CENTRAL, &jsonNBA.East),
			Atlantic:  convertToTeam(&jsonNBA.Atlantic, EAST, ATLANTIC, &jsonNBA.East),
			Total:     convertToTeam(&jsonNBA.East, EAST, 0, &emptyTeams),
		},

		West: West{
			Pacific:   convertToTeam(&jsonNBA.Pacific, WEST, PACIFIC, &jsonNBA.West),
			WestNorth: convertToTeam(&jsonNBA.WestNorth, WEST, WESTNORTH, &jsonNBA.West),
			WestSouth: convertToTeam(&jsonNBA.WestSouth, WEST, WESTSOUTH, &jsonNBA.West),
			Total:     convertToTeam(&jsonNBA.West, WEST, 0, &emptyTeams),
		},
	}
	result := common.ParseResult{Result: nba}

	for _, v := range nba.East.Total {
		v.Link = fmt.Sprintf(teamLinkTemp, v.TeamId)
		result.Requests = append(result.Requests, common.Request{
			Url:        v.Link,
			ParserFunc: parseRosters,
		})
	}

	for _, v := range nba.West.Total {
		v.Link = fmt.Sprintf(teamLinkTemp, v.TeamId)
		result.Requests = append(result.Requests, common.Request{
			Url:        v.Link,
			ParserFunc: parseRosters,
		})
	}

	return result
}

func findTemRank(team *JsonTeam, total *[]JsonTeam) int {
	for _, i := range *total {
		if i.TeamId == team.TeamId {
			rank, err := strconv.Atoi(i.Serial)
			common.PanicErr(err)
			return rank
		}
	}
	rank, err := strconv.Atoi(team.Serial)
	common.PanicErr(err)
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

func parseJson(content []byte) JsonNBA {
	var tempMap []interface{}
	err := json.Unmarshal(content, &tempMap)
	common.PanicErr(err)

	marshal, err := json.Marshal(tempMap[1])
	common.PanicErr(err)

	jsonNBA := JsonNBA{}
	err = json.Unmarshal(marshal, &jsonNBA)
	common.PanicErr(err)
	return jsonNBA
}

func convertToTeam(source *[]JsonTeam, areaId, divId int, total *[]JsonTeam) []Team {
	target := make([]Team, len(*source))
	for i, v := range *source {
		wins, err := strconv.Atoi(v.Wins)
		common.PanicErr(err)

		losses, err := strconv.Atoi(v.Losses)
		common.PanicErr(err)

		winingPercentage, err := strconv.Atoi(v.WiningPercentage)
		common.PanicErr(err)

		divRank, err := strconv.Atoi(v.DivRank)
		common.PanicErr(err)

		rank := findTemRank(&v, total)

		target[i] = Team{
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
