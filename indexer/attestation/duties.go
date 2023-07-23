package attestation

import (
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/config"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/db"
)

type Attestation struct {
	db              *db.Database
	chainURL        string
	dutiesStopChan  chan bool
	ValidatorsIndex []string
}

func New(db *db.Database, dutiesStopChan chan bool) *Attestation {
	return &Attestation{
		db:             db,
		chainURL:       config.ChainUrl,
		dutiesStopChan: dutiesStopChan,
	}
}
