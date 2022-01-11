package models

// Summary fundraising information for a poitician.
type CandidateSummary struct {
	CandidateName string  `json:"cand_name"`
	Cid           string  `json:"cid"` // CRP ID
	Cycle         int     `json:"cycle,string"`
	State         string  `json:"state"`                // Two-character abbreviation
	Party         string  `json:"party"`                // D, R, 3, L, U for Dem, Repub, 3rd party, Libertarian, Unknown
	Chamber       string  `json:"chamber"`              // S, H, D or blank
	FirstElected  int     `json:"first_elected,string"` // For members only, year first elected to current office
	NextElection  int     `json:"next_election,string"` // For members only, year of next election
	Total         float64 `json:"total,string"`         // Total receipts reported by candidate
	Spent         float64 `json:"spent,string"`         // Total expenditures reported by candidate
	CashOnHand    float64 `json:"cash_on_hand,string"`
	Debt          float64 `json:"debt,string"`
	Origin        string  `json:"origin"`       // Name for attribution
	Source        string  `json:"source"`       // Link to source data on OpenSecrets.org
	LastUpdated   string  `json:"last_updated"` // Date of candidate's last filed report (MM/DD/YYYY)
}
