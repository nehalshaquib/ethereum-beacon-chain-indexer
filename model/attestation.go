package model

type AttestationDuties struct {
	DependentRoot       string                  `json:"dependent_root"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
	Data                []AttestationDutiesData `json:"data"`
}

type AttestationDutiesData struct {
	PublicKey               string `json:"pubkey"`
	ValidatorIndex          string    `json:"validator_index"`
	CommitteesAtSlot        string    `json:"committees_at_slot"`
	CommitteeIndex          string    `json:"committee_index"`
	CommitteeLength         string    `json:"committee_length"`
	ValidatorCommitteeIndex string    `json:"validator_committee_index"`
	Slot                    string    `json:"slot"`
}

type BlockAttestationDetails struct {
	ExecutionOptimistic bool                   `json:"execution_optimistic"`
	Finalized           bool                   `json:"finalized"`
	Data                []BlockAttestationData `json:"data"`
}
type BlockAttestationData struct {
	AggregationBits string `json:"aggregation_bits"`
	Data            Data   `json:"data"`
	Signature       string `json:"signature"`
}

type Data struct {
	Slot            string        `json:"slot"`
	Index           string        `json:"index"`
	BeaconBlockRoot string     `json:"beacon_block_root"`
	Source          Checkpoint `json:"source"`
	Target          Checkpoint `json:"target"`
}

type Checkpoint struct {
	Epoch string    `json:"epoch"`
	Root  string `json:"root"`
}
