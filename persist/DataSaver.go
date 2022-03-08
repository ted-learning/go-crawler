package persist

import (
	"github.com/olivere/elastic/v7"
	"go-crawler/common"
	"go-crawler/parser"
	"golang.org/x/net/context"
	"log"
)

func DataSaver() chan interface{} {
	dataChan := make(chan interface{})
	client, err := elastic.NewClient(elastic.SetSniff(false))
	common.PanicErr(err)

	go func() {
		for {
			result := <-dataChan
			switch data := result.(type) {
			case parser.NBA:
				log.Println("East:")
				for _, v := range data.East.Total {
					saveTeam(client, v)
				}
				log.Println("West:")
				for _, v := range data.West.Total {
					saveTeam(client, v)
				}
			case []parser.Player:
				savePlayers(client, data)
			case parser.Stats:
				saveStats(client, data)
			}
		}
	}()
	return dataChan
}

func saveTeam(client *elastic.Client, v parser.Team) {
	re, err := client.Index().
		Index("team").
		Id(v.TeamId).
		BodyJson(v).
		Do(context.Background())
	common.PanicErr(err)
	log.Println(re)
}

func savePlayers(client *elastic.Client, players []parser.Player) {
	for _, v := range players {
		re, err := client.Index().
			Index("player").
			Id(v.PlayerId).
			BodyJson(v).
			Do(context.Background())
		common.PanicErr(err)
		log.Println(re)
	}
}

func saveStats(client *elastic.Client, v parser.Stats) {
	re, err := client.Index().
		Index("stats").
		Id(v.PlayerId).
		BodyJson(v).
		Do(context.Background())
	common.PanicErr(err)
	log.Println(re)
}
