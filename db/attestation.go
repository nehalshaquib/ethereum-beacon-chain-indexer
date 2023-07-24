package db

import (
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
)

func (db *Database) UpdateAttestationDuties(duties []model.AttestationDuties) (err error) {
	// Start a new transaction
	tx, err := db.db.Begin()
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			log.Println("unable to rollback")
		}
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Create a temporary table
	_, err = tx.Exec(`CREATE TEMP TABLE temp_attestation_duties (LIKE attestation_duties INCLUDING DEFAULTS) ON COMMIT DROP;`)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			log.Println("unable to rollback")
		}
		return fmt.Errorf("failed to create temporary table: %w", err)
	}

	// Prepare COPY statement
	stmt, err := tx.Prepare(pq.CopyIn("temp_attestation_duties", "pubkey", "validator_index", "committees_at_slot", "committee_index", "committee_length", "validator_committee_index", "slot"))
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			log.Println("err rollback: ", err1)
		}
		return fmt.Errorf("failed to prepare COPY statement: %w", err)
	}

	// Load data into the temporary table
	for _, duty := range duties {
		_, err = stmt.Exec(duty.PublicKey, duty.ValidatorIndex, duty.CommitteesAtSlot, duty.CommitteeIndex, duty.CommitteeLength, duty.ValidatorCommitteeIndex, duty.Slot)
		if err != nil {
			err1 := tx.Rollback()
			if err1 != nil {
				log.Println("err rollback: ", err1)
			}
			return fmt.Errorf("failed to exec COPY statement: %w", err)
		}
	}

	// Close the statement
	err = stmt.Close()
	if err != nil {
		err1 := tx.Rollback() // Ignore rollback error
		if err1 != nil {
			log.Println("err rollback: ", err1)
		}
		return fmt.Errorf("failed to close COPY statement: %w", err)
	}

	// Perform an INSERT
	_, err = tx.Exec(`INSERT INTO attestation_duties SELECT * FROM temp_attestation_duties;`)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			log.Println("err rollback: ", err1)
		}
		return fmt.Errorf("failed to insert/update validators: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) TruncateAttestationDuties() error {
	_, err := db.db.Exec("TRUNCATE TABLE attestation_duties;")
	if err != nil {
		return fmt.Errorf("failed to truncate attestation_duties table: %w", err)
	}
	return nil
}

func (db *Database) AddAttestations(response model.BlockAttestationResponse) error {
	for _, attestation := range response.Data {
		_, err := db.db.Exec(`INSERT INTO attestation_details (finalized, aggregation_bits, slot, index, beacon_block_root, source_epoch, source_root, target_epoch, target_root, signature, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			response.Finalized, attestation.AggregationBits, attestation.Data.Slot, attestation.Data.Index, attestation.Data.BeaconBlockRoot, attestation.Data.Source.Epoch, attestation.Data.Source.Root, attestation.Data.Target.Epoch, attestation.Data.Target.Root, attestation.Signature, time.Now())
		if err != nil {
			return fmt.Errorf("failed to add attestation: %w", err)
		}
	}

	return nil
}

func (db *Database) TruncateAttestationDetails() error {
	_, err := db.db.Exec("TRUNCATE TABLE attestation_details;")
	if err != nil {
		return fmt.Errorf("failed to truncate attestation_details table: %w", err)
	}
	return nil
}
