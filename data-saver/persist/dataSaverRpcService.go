package persist

import (
	"data-saver/common"
	"github.com/olivere/elastic/v7"
)

type DataSaverRpcService struct {
	Client *elastic.Client
}

func (s *DataSaverRpcService) SaveNBA(data common.NBA, result *string) error {
	err := SaveNBA(s.Client, data)
	return response(err, result)
}

func (s *DataSaverRpcService) SaveTeam(data common.Team, result *string) error {
	err := SaveTeam(s.Client, data)
	return response(err, result)
}

func (s *DataSaverRpcService) SavePlayers(data []common.Player, result *string) error {
	err := SavePlayers(s.Client, data)
	return response(err, result)
}

func (s *DataSaverRpcService) SaveStats(data common.Stats, result *string) error {
	err := SaveStats(s.Client, data)
	return response(err, result)
}

func response(err error, result *string) error {
	if err == nil {
		*result = "ok"
	}
	return err
}
