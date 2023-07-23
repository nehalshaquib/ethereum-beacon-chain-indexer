package util

import (
	"encoding/json"
	"net/http"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
)

func GetCheckPoints() (model.FinalityCheckpoints, error) {
	var checkpoints model.FinalityCheckpoints
	resp, err := http.Get(config.ChainUrl + "/eth/v1/beacon/states/head/finality_checkpoints")
	if err != nil {
		return checkpoints, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&checkpoints); err != nil {
		return checkpoints, err
	}

	return checkpoints, nil
}
