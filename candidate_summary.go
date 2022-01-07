package opensecrets

type CandidateSummary struct {
	CandidateName string  `json:"cand_name"`
	Cid           string  `json:"cid"`
	Cycle         int     `json:"cycle,string"`
	State         string  `json:"state"`
	Party         string  `json:"party"`
	Chamber       string  `json:"chamber"`
	FirstElected  int     `json:"first_elected,string"`
	NextElection  int     `json:"next_election,string"`
	Total         float64 `json:"total,string"`
	Spent         float64 `json:"spent,string"`
	CashOnHand    float64 `json:"cash_on_hand,string"`
	Debt          float64 `json:"debt,string"`
	Origin        string  `json:"origin"`
	Source        string  `json:"source"`
	LastUpdated   string  `json:"last_updated"`
}

type candidateSummaryWrapper struct {
	Attributes CandidateSummary `json:"@attributes"`
}

type candidateSummaryResponse struct {
	Wrapper candidateSummaryWrapper `json:"summary"`
}

type candidateSummaryResponseWrapper struct {
	Response candidateSummaryResponse `json:"response"`
}
