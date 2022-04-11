package client

import (
	"crawler-engine/common"
	"crawler-engine/distribute/rpc"
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
			case common.NBA:
				err := rpc.Call("DataSaverRpcService.SaveNBA", data, &result)
				common.PanicErr(err)
				var teams []common.Team
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
			case []common.Player:
				err := rpc.Call("DataSaverRpcService.SavePlayers", data, &result)
				common.PanicErr(err)
			case common.Stats:
				err := rpc.Call("DataSaverRpcService.SaveStats", data, &result)
				common.PanicErr(err)
			}
		}
	}()
	return dataChan
}
