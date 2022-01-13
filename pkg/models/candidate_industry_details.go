package models

// Total contributed to a candidate from a specific industry. Senate data reflects 2-year totals.
type CandidateIndustryDetails struct {
	CandidateName string  `json:"cand_name"`
	Cid           string  `json:"cid"` // CRP ID
	Cycle         int     `json:"cycle,string"`
	Industry      string  `json:"industry"`
	Chamber       string  `json:"chamber"`       // H or S for House or Senate
	Party         string  `json:"party"`         // D, R, 3, L, U for Dem, Repub, 3rd party, Libertarian, Unknown
	State         string  `json:"state"`         // Full state name
	Total         float64 `json:"total,string"`  // Total from all itemized sources
	Pacs          float64 `json:"pacs,string"`   // Total PAC contributions
	Individuals   float64 `json:"indivs,string"` // Total individual contributions
	Rank          int     `json:"rank,string"`   // Rank within chamber for this member
	Origin        string  `json:"origin"`        // Attribution to display
	Source        string  `json:"source"`        // Link to CRP data
	LastUpdated   string  `json:"last_updated"`  // Date data was last retrieved from government sources (MM/DD/YYYY)
}
