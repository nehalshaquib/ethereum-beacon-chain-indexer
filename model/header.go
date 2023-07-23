package model

type BlockHeader struct {
	ExecutionOptimistic bool            `json:"execution_optimistic"`
	Finalized           bool            `json:"finalized"`
	Data                BlockHeaderData `json:"data"`
}

type BlockHeaderData struct {
	Root      string `json:"root"`
	Canonical bool   `json:"canonical"`
	Header    Header `json:"header"`
}

type Header struct {
	Message   Message `json:"message"`
	Signature string  `json:"signature"`
}
type Message struct {
	Slot          string `json:"slot"`
	ProposerIndex string `json:"proposer_index"`
	ParentRoot    string `json:"parent_root"`
	StateRoot     string `json:"state_root"`
	BodyRoot      string `json:"body_root"`
}

type FinalityCheckpoints struct {
	Data struct {
		CurrentJustified struct {
			Epoch string `json:"epoch"`
		} `json:"current_justified"`
		Finalized struct {
			Epoch string `json:"epoch"`
		} `json:"finalized"`
	} `json:"data"`
}
