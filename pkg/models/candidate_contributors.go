/*
Package model provides data types representing responses from the OpenSecrets REST API.
Some of the fields on these types, like Cid, are explained in detail in the OpenSecrets OpenData
User's Guide: https://www.opensecrets.org/resources/datadictionary/UserGuide.pdf
*/
package models

// Details on top contributors to a candidate for a House or Senate seat, or a sitting member of Congress.
// These are 6-year running numbers for senators/Senate candidates, 2-year numbers for representatives/House candidates.
type CandidateContributorSummary struct {
	CandidateName string `json:"cand_name"`
	Cid           string `json:"cid"`          // CRP ID
	Cycle         int    `json:"cycle,string"` // Cycle year of data being returned
	Origin        string `json:"origin"`       // Attribution to display
	Source        string `json:"source"`       // Link to CRP data
	Notice        string `json:"notice"`       // Required explanatory text - must be displayed with published data
	Contributors  []CandidateContributor
}

// A contributor to a candidate.
type CandidateContributor struct {
	OrganizationName string  `json:"org_name"`
	Total            float64 `json:"total,string"`  // Total from all itemized sources
	Pacs             float64 `json:"pacs,string"`   // Total PAC contributions
	Individuals      float64 `json:"indivs,string"` // Total individual contributions
}
