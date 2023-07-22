package model

type AttestationDuties struct {
	DependentRoot       string                  `json:"dependent_root"`
	ExecutionOptimistic bool                    `json:"execution_optimistic"`
	Data                []AttestationDutiesData `json:"data"`
}

type AttestationDutiesData struct {
	PublicKey               string `json:"pubkey"`
	ValidatorIndex          int    `json:"validator_index"`
	CommitteesAtSlot        int    `json:"committees_at_slot"`
	CommitteeIndex          int    `json:"committee_index"`
	CommitteeLength         int    `json:"committee_length"`
	ValidatorCommitteeIndex int    `json:"validator_committee_index"`
	Slot                    int    `json:"slot"`
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
	Slot            int        `json:"slot"`
	Index           int        `json:"index"`
	BeaconBlockRoot string     `json:"beacon_block_root"`
	Source          Checkpoint `json:"source"`
	Target          Checkpoint `json:"target"`
}

type Checkpoint struct {
	Epoch int    `json:"epoch"`
	Root  string `json:"root"`
}
