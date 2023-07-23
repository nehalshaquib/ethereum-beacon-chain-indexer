package db

import (
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/model"
)

func (db *Database) AddOrUpdateValidators(validators []model.Validator) (err error) {
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
	_, err = tx.Exec(`CREATE TEMP TABLE temp_validators (LIKE validators INCLUDING DEFAULTS) ON COMMIT DROP;`)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			log.Println("unable to rollback")
		}
		return fmt.Errorf("failed to create temporary table: %w", err)
	}

	// Prepare COPY statement
	stmt, err := tx.Prepare(pq.CopyIn("temp_validators", "index", "balance", "status", "pubkey", "withdrawal_credentials", "effective_balance", "slashed", "activation_eligibility_epoch", "activation_epoch", "exit_epoch", "withdrawable_epoch"))
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			log.Println("err rollback: ", err1)
		}
		return fmt.Errorf("failed to prepare COPY statement: %w", err)
	}

	// Load data into the temporary table
	for _, validator := range validators {
		_, err = stmt.Exec(validator.Index, validator.Balance, validator.Status, validator.Validator.Pubkey, validator.Validator.WithdrawalCredentials, validator.Validator.EffectiveBalance, validator.Validator.Slashed, validator.Validator.ActivationEligibilityEpoch, validator.Validator.ActivationEpoch, validator.Validator.ExitEpoch, validator.Validator.WithdrawableEpoch)
		if err != nil {
			log.Println("ERROR in Load: ", err)
			err1 := tx.Rollback()
			if err1 != nil {
				log.Println("err rollback: ", err1)
			}
			return fmt.Errorf("failed to exec COPY statement: %w", err)
		}
		log.Println("Done for: ", validator.Index)
	}

	// Execute the statement
	_, err = stmt.Exec()
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			log.Println("err rollback: ", err1)
		}
		return fmt.Errorf("failed to execute COPY statement: %w", err)
	}

	// Perform an INSERT
	_, err = tx.Exec(`INSERT INTO validators SELECT * FROM temp_validators ON CONFLICT (index) DO UPDATE SET balance = EXCLUDED.balance, status = EXCLUDED.status, pubkey = EXCLUDED.pubkey, withdrawal_credentials = EXCLUDED.withdrawal_credentials, effective_balance = EXCLUDED.effective_balance, slashed = EXCLUDED.slashed, activation_eligibility_epoch = EXCLUDED.activation_eligibility_epoch, activation_epoch = EXCLUDED.activation_epoch, exit_epoch = EXCLUDED.exit_epoch, withdrawable_epoch = EXCLUDED.withdrawable_epoch;`)
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
