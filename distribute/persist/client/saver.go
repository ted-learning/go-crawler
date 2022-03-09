package client

import (
	"go-crawler/common"
	rpcsupport "go-crawler/distribute/rpc"
	"go-crawler/parser"
)

func DataSaver(host string) chan interface{} {
	dataChan := make(chan interface{})
	rpc, err := rpcsupport.NewClient(host)
	common.PanicErr(err)

	go func() {
		for {
			data := <-dataChan
			result := ""
			switch data := data.(type) {
			case parser.NBA:
				err := rpc.Call("DataSaverRpcService.SaveNBA", data, &result)
				common.PanicErr(err)
				var teams []parser.Team
				teams = append(teams, data.East.EastSouth...)
				teams = append(teams, data.East.Atlantic...)
				teams = append(teams, data.East.Central...)
				teams = append(teams, data.West.Pacific...)
				teams = append(teams, data.West.WestSouth...)
				teams = append(teams, data.West.WestNorth...)
				for _, v := range teams {
					err := rpc.Call("DataSaverRpcService.SaveTeam", v, &result)
					common.PanicErr(err)
				}
			case []parser.Player:
				err := rpc.Call("DataSaverRpcService.SavePlayers", data, &result)
				common.PanicErr(err)
			case parser.Stats:
				err := rpc.Call("DataSaverRpcService.SaveStats", data, &result)
				common.PanicErr(err)
			}
		}
	}()
	return dataChan
}
