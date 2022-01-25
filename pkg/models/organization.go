package models

// Result of a search by organization name
type OrganizationSearchResult struct {
	Id   string `json:"orgid"`   // CRP org ID
	Name string `json:"orgname"` // Standardized org name
}

// Summary of an organization's fundraising information
type OrganizationSummary struct {
	Id                                  string  `json:"orgid"` // CPR org ID
	Cycle                               string  `json:"cycle"`
	Name                                string  `json:"orgname"`       // Standardized org name
	TotalContributions                  float64 `json:"total,string"`  // Total contributions (FEC and IRS)
	PacContributions                    float64 `json:"pacs,string"`   // Total from organization's PACs
	IndividualContributions             float64 `json:"indivs,string"` // Total from individuals
	Soft                                float64 `json:"soft,string"`   // Total soft money
	TotalFrom527Organizations           float64 `json:"tot527,string"`
	TotalToDemocrats                    float64 `json:"dems,string"`
	TotalToRepublicans                  float64 `json:"repubs,string"`
	TotalSpentLobyying                  float64 `json:"lobbying,string"`
	TotalSpentOnIndependentExpenditures float64 `json:"outside,string"`
	MembersInvested                     int     `json:"mems_invested,string"` // Number of members invested in the organization
	TotalGaveToPacs                     float64 `json:"gave_to_pac,string"`
	TotalGaveToPartyCommittees          float64 `json:"gave_to_party,string"`
	TotalGaveTo527Organizations         float64 `json:"gave_to_527,string"`
	TotalGaveToCandidates               float64 `json:"gave_to_cand,string"`
	Source                              string  `json:"source"` // Link to CRP data
}
