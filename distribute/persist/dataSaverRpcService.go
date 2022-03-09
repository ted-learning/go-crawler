package persist

import (
	"github.com/olivere/elastic/v7"
	"go-crawler/parser"
	"go-crawler/persist"
)

type DataSaverRpcService struct {
	Client *elastic.Client
}

func (s *DataSaverRpcService) SaveNBA(data parser.NBA, result *string) error {
	err := persist.SaveNBA(s.Client, data)
	return response(err, result)
}

func (s *DataSaverRpcService) SaveTeam(data parser.Team, result *string) error {
	err := persist.SaveTeam(s.Client, data)
	return response(err, result)
}

func (s *DataSaverRpcService) SavePlayers(data []parser.Player, result *string) error {
	err := persist.SavePlayers(s.Client, data)
	return response(err, result)
}

func (s *DataSaverRpcService) SaveStats(data parser.Stats, result *string) error {
	err := persist.SaveStats(s.Client, data)
	return response(err, result)
}

func response(err error, result *string) error {
	if err == nil {
		*result = "ok"
	}
	return err
}
