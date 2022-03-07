package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
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
	Serial           int
	Link             string
}

func findNBATeams(content []byte, decoder *encoding.Decoder) {
	//find ajax url
	compile := regexp.MustCompile("(n\\.ajax\\({url:\"//)([a-zA-Z0-9&=_/.?]+)(\")")
	submatch := compile.FindSubmatch(content)
	//request to get teams json data
	response, err := http.Get("https://" + string(submatch[2]))
	panicErr(err)
	jsonData, err := ioutil.ReadAll(transform.NewReader(response.Body, decoder))
	panicErr(err)
	//find team details link format
	link := findTeamsLink(content)

	//parse to NBA structs
	NBA := parseNBATeams(jsonData, link)

	for i, v := range NBA.East.Total {
		fmt.Printf("%d %v\n", i, v)
	}

	for i, v := range NBA.West.Total {
		fmt.Printf("%d %v\n", i, v)
	}
}

func parseNBATeams(jsonData []byte, teamLink string) NBA {
	tempMap := make([]interface{}, 3)
	err := json.Unmarshal(jsonData, &tempMap)
	panicErr(err)

	marshal, err := json.Marshal(tempMap[1])
	panicErr(err)

	jsonNBA := JsonNBA{}
	err = json.Unmarshal(marshal, &jsonNBA)
	panicErr(err)

	east := convertToTeam(jsonNBA.East, teamLink)
	west := convertToTeam(jsonNBA.West, teamLink)

	eastSouth := convertToTeam(jsonNBA.EastSouth, teamLink)
	central := convertToTeam(jsonNBA.Central, teamLink)
	atlantic := convertToTeam(jsonNBA.Atlantic, teamLink)

	pacific := convertToTeam(jsonNBA.Pacific, teamLink)
	westNorth := convertToTeam(jsonNBA.WestNorth, teamLink)
	westSouth := convertToTeam(jsonNBA.WestSouth, teamLink)

	return NBA{
		East: East{
			EastSouth: eastSouth, //东南赛区
			Central:   central,   //中部赛区
			Atlantic:  atlantic,  //大西洋赛区
			Total:     east,
		},

		West: West{
			Pacific:   pacific,   //太平洋赛区
			WestNorth: westNorth, //西北赛区
			WestSouth: westSouth, //西南赛区
			Total:     west,
		},
	}
}

func findTeamsLink(data []byte) string {
	compile := regexp.MustCompile("(https://[a-zA-Z0-9&=_/.?]+)(['+s.teamId]*)([&cid=0-9]*)")
	s := compile.FindSubmatch(data)
	return string(s[1]) + "%s" + string(s[3])
}

func convertToTeam(source []JsonTeam, link string) []Team {
	target := make([]Team, len(source))
	for i, v := range source {
		wins, err := strconv.Atoi(v.Wins)
		panicErr(err)

		losses, err := strconv.Atoi(v.Losses)
		panicErr(err)

		winingPercentage, err := strconv.Atoi(v.WiningPercentage)
		panicErr(err)

		rank, err := strconv.Atoi(v.DivRank)
		panicErr(err)

		serial, err := strconv.Atoi(v.Serial)
		panicErr(err)

		target[i] = Team{
			Badge:            v.Badge,
			Name:             v.Name,
			TeamId:           v.TeamId,
			Wins:             wins,
			Losses:           losses,
			WiningPercentage: winingPercentage,
			DivRank:          rank,
			Serial:           serial,
			Link:             fmt.Sprintf(link, v.TeamId),
		}
	}

	sort.Slice(target, func(i, j int) bool {
		return target[i].Serial < target[j].Serial
	})

	return target
}
