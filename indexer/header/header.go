package header

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/db"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
)

type Header struct {
	db            *db.Database
	chainURL      string
	headerErrChan chan error
}

func New(db *db.Database, headerErrChan chan error) *Header {
	return &Header{
		db:            db,
		chainURL:      config.ChainUrl,
		headerErrChan: headerErrChan,
	}
}

func (h *Header) getNewSlot(blockID string) (*model.BlockHeader, error) {

	resp, err := http.Get(fmt.Sprintf("%v/eth/v1/beacon/headers/%v", h.chainURL, blockID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var header model.BlockHeader
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &header)
	if err != nil {
		return nil, err
	}
	return &header, nil
}

func (h *Header) getLatestBlockID() (string, error) {
	block, err := h.getNewSlot("head")
	if err != nil {
		return "", err
	}
	return block.Data.Header.Message.Slot, nil
}

func (h *Header) SlotHeaderIndexer() {
	blockID, err := h.getLatestBlockID()
	if err != nil {
		h.headerErrChan <- err
		return
	}

	for {
		header, err := h.getNewSlot(blockID)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		_, err = h.db.AddSlot(header)
		if err != nil {
			h.headerErrChan <- err
		}
		log.Println("get Latest block done in main")
		// Increment the block ID for the next slot
		blockIDInt, err := strconv.Atoi(blockID)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		blockID = strconv.Itoa(blockIDInt + 1)

		// Wait for the next slot
		time.Sleep(12 * time.Second)
	}
}
