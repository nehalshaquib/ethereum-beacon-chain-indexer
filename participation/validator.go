package participation

import (
	"log"

	"github.com/nehalshaquib/ethereum-beacon-chain-indexer.git/db"
)

type Participation struct {
	db *db.Database
}

func New(db *db.Database) *Participation {
	return &Participation{
		db: db,
	}
}
func (p *Participation) GetParticipationRate(validatorIndex string) (float64, error) {
	// Calculate the total number of attestations
	totalAttestations, err := p.db.ValidatorSupposedAttestations(validatorIndex)
	if err != nil {
		return 0, err
	}
	log.Println("Total Attestations:", totalAttestations)
	// Count the number of attestations made by the validator
	madeAttestations, err := p.db.ValidatorAttestationsMade(validatorIndex)
	if err != nil {
		return 0, err
	}

	log.Println("made Attestations:", madeAttestations)
	// Calculate the participation rate
	participationRate := 1 - float64(totalAttestations-madeAttestations)/float64(totalAttestations)

	return participationRate, nil
}
