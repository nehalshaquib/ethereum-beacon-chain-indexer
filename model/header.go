package model

type BlockHeader struct {
	ExecutionOptimistic bool            `json:"execution_optimistic"`
	Finalized           bool            `json:"finalized"`
	Data                BlockHeaderData `json:"data"`
}

type BlockHeaderData struct {
	Root      string `json:"root" db:"root"`
	Canonical bool   `json:"canonical" db:"canonical"`
	Header    struct {
		Message struct {
			Slot          string `json:"slot" db:"slot"`
			ProposerIndex string `json:"proposer_index" db:"proposer_index"`
			ParentRoot    string `json:"parent_root" db:"parent_root"`
			StateRoot     string `json:"state_root" db:"state_root"`
			BodyRoot      string `json:"body_root" db:"body_root"`
		} `json:"message"`
		Signature string `json:"signature"`
	} `json:"header"`
}
