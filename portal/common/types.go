package common

type Request struct {
	Url        string
	Context    Context
	ParserFunc func([]byte, Context) ParseResult
}

type ParseResult struct {
	Requests []Request
	Result   interface{}
}

type Context map[string]interface{}

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
