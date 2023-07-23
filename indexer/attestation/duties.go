package attestation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/db"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/util"
)

type Attestation struct {
	db                     *db.Database
	chainURL               string
	dutiesStopChan         chan bool
	validatorsIndex        []string
	validatorUpdateStarted chan bool
}

func New(db *db.Database, dutiesStopChan chan bool, validatorUpdateStarted chan bool) *Attestation {
	return &Attestation{
		db:                     db,
		chainURL:               config.ChainUrl,
		dutiesStopChan:         dutiesStopChan,
		validatorUpdateStarted: validatorUpdateStarted,
	}
}

func (a *Attestation) getAttestationDuties(epoch string, validatorIndices []string) ([]model.AttestationDuties, error) {
	// Convert the validatorIndices slice to JSON
	jsonIndices, err := json.Marshal(validatorIndices)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal validator indices: %w", err)
	}

	// Create a new HTTP request
	reqUrl := fmt.Sprintf("%s/eth/v1/validator/duties/attester/%s", config.ChainUrl, epoch)
	log.Println("Request URL attestation: ", reqUrl)
	log.Println("Request BODY: ", len(validatorIndices))
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonIndices))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	log.Println("Resp Status: ", resp.StatusCode)
	dec := json.NewDecoder(resp.Body)
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
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	log.Println("TOKEN: ", t)
	// Decode the response
	var duties []model.AttestationDuties
	for dec.More() {
		var duty model.AttestationDuties
		err := dec.Decode(&duty)
		if err != nil {
			return nil, err
		}
		log.Println("attestation for: ", duty.ValidatorIndex)
		duties = append(duties, duty)
	}

	return duties, nil
}

func (a *Attestation) UpdateAttestationDuties() error {
	<-a.validatorUpdateStarted
	fmt.Println("UpdateAttestationDuties Started")
	epochsPassed := 0
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
			epochsPassed++
			if epochsPassed > 5 {
				err = a.db.TruncateAttestationDuties()
				if err != nil {
					log.Println("TruncateAttestationDuties:", err)
				}
			}
			log.Println("getting attestaton duties")

			duties, err := a.getAttestationDuties(currentFinalizedEpoch, a.validatorsIndex)
			if err != nil {
				log.Println("getAttestationDuties:", err)
			}

			log.Println("updating attestaton duties")
			// Add or update the validators in the database
			err = a.db.UpdateAttestationDuties(duties)
			if err != nil {
				log.Println("Failed to add or update attestaton duties:", err)
			}

			log.Println("attestaton duties updated")
			// Update the previous epoch
			previousFinalizedEpoch = currentFinalizedEpoch
		}

		log.Println("waiting till next epoch")
		// Sleep until the next epochs
		time.Sleep(time.Duration(32*12) * time.Second)
	}

}
