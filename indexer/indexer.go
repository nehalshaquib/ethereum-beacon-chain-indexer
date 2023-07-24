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
	headerStopChan         chan bool
	IndexerStopChan        chan bool
	validatorUpdateStarted chan bool
}

func New(db *db.Database) *Indexer {
	return &Indexer{
		db:                     db,
		headerStopChan:         make(chan bool),
		IndexerStopChan:        make(chan bool),
		validatorUpdateStarted: make(chan bool),
	}
}

func (i *Indexer) IndexHeader() {
	h := header.New(i.db, i.headerStopChan)
	go func() {
		h.SlotHeaderIndexer()
	}()
}

func (i *Indexer) StartAttestations() {
	a := attestation.New(i.db, i.validatorUpdateStarted)
	go func() {
		err := a.UpdateActiveValidators()
		if err != nil {
			log.Println("ERR: ", err)
		}
	}()

	go func() {
		err := a.UpdateAttestationDuties()
		if err != nil {
			log.Println("ERR: ", err)
		}
	}()

	go a.StartBlockAttestationDetails()
	<-i.IndexerStopChan
}
