package main

import (
	"log"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/db"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/indexer"
)

func main() {
	err := config.Config()
	if err != nil {
		log.Println("ERR: config - ", err)
	}

	db := db.NewDatabase()

	indexer := indexer.New(db)
	indexer.IndexHeader()
	indexer.StartAttestations()
}
