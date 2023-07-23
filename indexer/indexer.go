package indexer

import (
	"log"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/db"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/indexer/attestation"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/indexer/header"
)

type Indexer struct {
	db *db.Database
	// chainUrl      string
	headerStopChan chan bool
	dutiesStopChan chan bool
}

func New(db *db.Database) *Indexer {
	return &Indexer{
		db:             db,
		headerStopChan: make(chan bool),
		dutiesStopChan: make(chan bool),
	}
}

func (i *Indexer) IndexHeader() {
	h := header.New(i.db, i.headerStopChan)
	go func() {
		h.SlotHeaderIndexer()
	}()
}

func (i *Indexer) UpdateValidators() {
	a := attestation.New(i.db, i.dutiesStopChan)
	go func() {
		err := a.UpdateActiveValidators()
		if err != nil {
			log.Println("ERR: ", err)
		}
	}()
	<-i.dutiesStopChan
}
