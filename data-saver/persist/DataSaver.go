package persist

import (
	"context"
	"data-saver/common"
	"github.com/olivere/elastic/v7"
	"log"
)

//goland:noinspection GoUnusedExportedFunction
func DataSaver() chan interface{} {
	dataChan := make(chan interface{})
	client, err := elastic.NewClient(elastic.SetSniff(false))
	common.PanicErr(err)

	go func() {
		for {
			result := <-dataChan
			switch data := result.(type) {
			case common.NBA:
				err := SaveNBA(client, data)
				common.PanicErr(err)
				var teams []common.Team
				teams = append(teams, data.East.EastSouth...)
				teams = append(teams, data.East.Atlantic...)
				teams = append(teams, data.East.Central...)
				teams = append(teams, data.West.Pacific...)
				teams = append(teams, data.West.WestSouth...)
				teams = append(teams, data.West.WestNorth...)
				for _, v := range teams {
					err := SaveTeam(client, v)
					common.PanicErr(err)
				}
			case []common.Player:
				err := SavePlayers(client, data)
				common.PanicErr(err)
			case common.Stats:
				err := SaveStats(client, data)
				common.PanicErr(err)
			}
		}
	}()
	return dataChan
}

func SaveNBA(client *elastic.Client, v common.NBA) error {
	re, err := client.Index().
		Index("nba").
		Id("nba").
		BodyJson(v).
		Do(context.Background())
	if err != nil {
		return err
	}
	log.Println(re)
	return nil
}

func SaveTeam(client *elastic.Client, v common.Team) error {
	re, err := client.Index().
		Index("team").
		Id(v.TeamId).
		BodyJson(v).
		Do(context.Background())
	if err != nil {
		return err
	}
	log.Println(re)
	return nil
}

func SavePlayers(client *elastic.Client, players []common.Player) error {
	for _, v := range players {
		re, err := client.Index().
			Index("player").
			Id(v.PlayerId).
			BodyJson(v).
			Do(context.Background())
		if err != nil {
			return err
		}
		log.Println(re)
	}
	return nil
}

func SaveStats(client *elastic.Client, v common.Stats) error {
	re, err := client.Index().
		Index("stats").
		Id(v.PlayerId).
		BodyJson(v).
		Do(context.Background())
	if err != nil {
		return err
	}
	log.Println(re)
	return nil
}
