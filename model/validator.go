package model

type Validator struct {
	Index     string `json:"index"`
	Balance   string `json:"balance"`
	Status    string `json:"status"`
	Validator struct {
		Pubkey                     string `json:"pubkey"`
		WithdrawalCredentials      string `json:"withdrawal_credentials"`
		EffectiveBalance           string `json:"effective_balance"`
		Slashed                    bool   `json:"slashed"`
		ActivationEligibilityEpoch string `json:"activation_eligibility_epoch"`
		ActivationEpoch            string `json:"activation_epoch"`
		ExitEpoch                  string `json:"exit_epoch"`
		WithdrawableEpoch          string `json:"withdrawable_epoch"`
	} `json:"validator"`
}

type ValidatorResponse struct {
	Finalized bool        `json:"finalized"`
	Data      []Validator `json:"data"`
}
