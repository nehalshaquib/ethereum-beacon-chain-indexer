package db

import (
	"fmt"
)

func (db *Database) ValidatorAttestationsMade(ValidatorIndex string) (int, error) {
	var madeAttestations int
	err := db.db.QueryRow(`SELECT COUNT(*) FROM attestation_details INNER JOIN attestation_duties ON attestation_details.slot = attestation_duties.slot AND attestation_details.index = attestation_duties.committee_index WHERE attestation_duties.validator_index = $1`, ValidatorIndex).Scan(&madeAttestations)
	if err != nil {
		return 0, fmt.Errorf("failed to count attestations: %w", err)
	}

	return madeAttestations, nil
}

func (db *Database) ValidatorSupposedAttestations(ValidatorIndex string) (int, error) {
	var supposedAttestations int
	err := db.db.QueryRow(`SELECT COUNT(*) FROM attestation_duties WHERE validator_index = $1 `, ValidatorIndex).Scan(&supposedAttestations)
	if err != nil {
		return 0, fmt.Errorf("failed to count total attestations: %w", err)
	}
	return supposedAttestations, nil
}
