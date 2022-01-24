package models

// Summary fundraising information for a congressional committee.
type CommitteeFundraisingDetails struct {
	CommitteeName  string `json:"committee_name"`
	Industry       string `json:"industry"` // The name of the industry
	CongressNumber int    `json:"congno,string"`
	Origin         string `json:"origin"`       // Attribution to display
	Source         string `json:"source"`       // Link to CRP data
	LastUpdated    string `json:"last_updated"` // Date data was last retrieved from government sources (MM/DD/YYYY)
	Members        []CommitteeMember
}

// Details on a member of a congressional committee
type CommitteeMember struct {
	Name        string  `json:"member_name"`
	Cid         string  `json:"cid"`           // CRP ID
	Party       string  `json:"party"`         // D, R, 3, L, U for Dem, Repub, 3rd party, Libertarian, Unknown
	State       string  `json:"state"`         // Full state name
	Total       float64 `json:"total,string"`  // Total from all itemized sources in the industry
	Pacs        float64 `json:"pacs,string"`   // Total PAC contributions from the industry
	Individuals float64 `json:"indivs,string"` // Total individual contributions from the industry
}
