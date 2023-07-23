package attestation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/util"
)

func (a *Attestation) getValidators() ([]model.Validator, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/states/head/validators?status=active", config.ChainUrl)
	log.Println("GetValidators URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get validators req: %w", err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	log.Println("POINT-1")
	// Read the opening bracket '['
	for {
		t, err := dec.Token()
		if err != nil {
			return nil, err
		}

		if t == "data" {
			break
		}
	}

	// Read the opening bracket '['
	_, err = dec.Token()
	if err != nil {
		return nil, err
	}

	var validators []model.Validator
	for dec.More() {
		var validator model.Validator
		// log.Println("POINT-3")
		err := dec.Decode(&validator)
		if err != nil {
			return nil, err
		}
		// log.Println("POINT-4")

		validators = append(validators, validator)
	}
	return validators, nil
}

func (a *Attestation) UpdateActiveValidators() error {
	previousFinalizedEpoch := "0"
	for {
		log.Println("updating checkpoints")
		// Get the current epoch
		checkpoint, err := util.GetCheckPoints()
		if err != nil {
			log.Println("Failed to get current epoch:", err)
		}

		currentFinalizedEpoch := checkpoint.Data.Finalized.Epoch
		log.Println("last finalized epoch: ", currentFinalizedEpoch)

		// Check if the epoch has changed
		if currentFinalizedEpoch != previousFinalizedEpoch {
			log.Println("getting validators data")

			// Fetch the validators for the current epoch
			validators, err := a.getValidators()
			if err != nil {
				log.Println("getValidators:", err)
			}

			log.Println("updating validators data")
			// Add or update the validators in the database
			err = a.db.AddOrUpdateValidators(validators)
			if err != nil {
				log.Println("Failed to add or update validators:", err)
			}

			log.Println("validators validators updated")
			// Update the previous epoch
			previousFinalizedEpoch = currentFinalizedEpoch
		}

		log.Println("waiting till next 5 epochs (32 minutes)")
		// Sleep until the next 5 epochs
		time.Sleep(time.Duration(32*12*5) * time.Second)
	}

}
