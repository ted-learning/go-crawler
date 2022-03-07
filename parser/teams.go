package parser

import (
	"encoding/json"
	"fmt"
	"go-crawler/common"
	"go-crawler/engine"
	"log"
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

//type JsonTeams []JsonTeam

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
	Serial           int
	Link             string
}

func parseTeams(content []byte, ctx engine.Context) engine.ParseResult {
	jsonNBA := parseJson(content)
	nba := NBA{
		East: East{
			EastSouth: convertToTeam(&jsonNBA.EastSouth), //东南赛区
			Central:   convertToTeam(&jsonNBA.Central),   //中部赛区
			Atlantic:  convertToTeam(&jsonNBA.Atlantic),  //大西洋赛区
			Total:     convertToTeam(&jsonNBA.East),
		},

		West: West{
			Pacific:   convertToTeam(&jsonNBA.Pacific),   //太平洋赛区
			WestNorth: convertToTeam(&jsonNBA.WestNorth), //西北赛区
			WestSouth: convertToTeam(&jsonNBA.WestSouth), //西南赛区
			Total:     convertToTeam(&jsonNBA.West),
		},
	}
	result := engine.ParseResult{Result: nba}

	linkFormat := ctx[linkFormatKey].(string)
	log.Println("East:")
	for i, v := range nba.East.Total {
		v.Link = fmt.Sprintf(linkFormat, v.TeamId)
		result.Requests = append(result.Requests, engine.Request{
			Url:        v.Link,
			ParserFunc: parseRosters,
		})
		log.Printf("%d %v\n", i, v)
	}
	log.Println("West:")
	for i, v := range nba.West.Total {
		v.Link = fmt.Sprintf(linkFormat, v.TeamId)
		result.Requests = append(result.Requests, engine.Request{
			Url:        v.Link,
			ParserFunc: parseRosters,
		})
		log.Printf("%d %v\n", i, v)
	}

	return result
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

func convertToTeam(source *[]JsonTeam) []Team {
	target := make([]Team, len(*source))
	for i, v := range *source {
		wins, err := strconv.Atoi(v.Wins)
		common.PanicErr(err)

		losses, err := strconv.Atoi(v.Losses)
		common.PanicErr(err)

		winingPercentage, err := strconv.Atoi(v.WiningPercentage)
		common.PanicErr(err)

		rank, err := strconv.Atoi(v.DivRank)
		common.PanicErr(err)

		serial, err := strconv.Atoi(v.Serial)
		common.PanicErr(err)

		target[i] = Team{
			Badge:            v.Badge,
			Name:             v.Name,
			TeamId:           v.TeamId,
			Wins:             wins,
			Losses:           losses,
			WiningPercentage: winingPercentage,
			DivRank:          rank,
			Serial:           serial,
		}
	}

	sort.Slice(target, func(i, j int) bool {
		return target[i].Serial < target[j].Serial
	})

	return target
}
