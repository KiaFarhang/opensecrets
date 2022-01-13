package models

// Details on the top industries contributing to a candidate for a House/Senate seat or a member of Congress.
// These are 6-year numbers for the Senate and 2-year numbers for the House.
type CandidateIndustriesSummary struct {
	CandidateName string `json:"cand_name"`
	Cid           string `json:"cid"` // CRP ID
	Cycle         int    `json:"cycle,string"`
	Origin        string `json:"origin"`       // Attribution to display
	Source        string `json:"source"`       // Link to CRP data
	LastUpdated   string `json:"last_updated"` // Date data was last retrieved from government sources (MM/DD/YYYY)
	Industries    []Industry
}

// An industry individuals/PACs belong to
type Industry struct {
	IndustryCode string  `json:"industry_code"` // CRP ID for the industry
	IndustryName string  `json:"industry_name"`
	Total        float64 `json:"total,string"`  // Total from all itemized sources
	Pacs         float64 `json:"pacs,string"`   // Total PAC contributions
	Individuals  float64 `json:"indivs,string"` // Total individual contributions
}
