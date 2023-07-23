package db

import (
	"database/sql"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
)

func (db *Database) AddSlot(header *model.BlockHeader) (sql.Result, error) {
	res, err := db.db.Exec(`INSERT INTO block_headers (slot, root, canonical, proposer_index, parent_root, state_root, body_root, signature, finalized) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		header.Data.Header.Message.Slot, header.Data.Root, header.Data.Canonical, header.Data.Header.Message.ProposerIndex, header.Data.Header.Message.ParentRoot, header.Data.Header.Message.StateRoot, header.Data.Header.Message.BodyRoot, header.Data.Header.Signature, header.Finalized)
	if err != nil {
		return nil, err
	}
	return res, nil
}
