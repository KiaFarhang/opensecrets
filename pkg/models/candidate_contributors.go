package models

type CandidateContributorSummary struct {
	CandidateName string `json:"cand_name"`
	Cid           string `json:"cid"`
	Cycle         int    `json:"cycle,string"`
	Origin        string `json:"origin"`
	Source        string `json:"source"`
	Notice        string `json:"notice"`
	Contributors  []CandidateContributor
}

type CandidateContributor struct {
	OrganizationName string  `json:"org_name"`
	Total            float64 `json:"total,string"`
	Pacs             float64 `json:"pacs,string"`
	Individuals      float64 `json:"indivs,string"`
}
