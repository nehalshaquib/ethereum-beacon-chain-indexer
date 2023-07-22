package indexer

import (
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/db"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/indexer/header"
)

type Indexer struct {
	db *db.Database
	// chainUrl      string
	headerErrChan chan error
}

func New(db *db.Database) *Indexer {
	return &Indexer{
		db:            db,
		headerErrChan: make(chan error),
	}
}

func (i *Indexer) IndexHeader() error {
	h := header.New(i.db, i.headerErrChan)
	go func() {
		h.SlotHeaderIndexer()
	}()
	return <-i.headerErrChan
}
