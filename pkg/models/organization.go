package models

// Result of a search by organization name
type OrganizationSearchResult struct {
	Id   string `json:"orgid"`   // CRP org ID
	Name string `json:"orgname"` // Standardized org name
}
