package attestation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
)

func (a *Attestation) getBlockAttestations(block string) (model.BlockAttestationResponse, error) {
	var blockAttestationsResponse model.BlockAttestationResponse

	url := fmt.Sprintf("%s/eth/v1/beacon/blocks/%s/attestations", config.ChainUrl, block)

	resp, err := http.Get(url)
	if err != nil {
		return blockAttestationsResponse, fmt.Errorf("failed to get block attestations: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return blockAttestationsResponse, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &blockAttestationsResponse)
	if err != nil {
		return blockAttestationsResponse, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return blockAttestationsResponse, nil
}

func (a *Attestation) StartBlockAttestationDetails() {
	currentBlock, previousBlock := "", ""
	slotsPassed := 0
	for {
		response, err := a.getBlockAttestations("head")
		if err != nil {
			log.Printf("Failed to get attestations for latest block: %v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}
		currentBlock = response.Data[0].Data.Slot

		if currentBlock != previousBlock {
			slotsPassed++
			if slotsPassed > 160 {
				log.Printf("truncating attestation details in slot %s as passed 160 slots", currentBlock)
				err = a.db.TruncateAttestationDetails()
				if err != nil {
					log.Println("TruncateAttestationDetails:", err)
				}
			}

			err = a.db.AddAttestations(response)
			if err != nil {
				log.Printf("Failed to add attestations for latest block: %v\n", err)
				time.Sleep(10 * time.Second)
				continue
			}
			log.Println("Added attestations for latest block")
		} else {
			log.Println("Slot not changed, skipping updating db")
		}
		previousBlock = currentBlock

		// Sleep for a while before fetching the next block
		time.Sleep(12 * time.Second)
	}
}
