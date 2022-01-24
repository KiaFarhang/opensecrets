package models

// Details about the sector total of a politician's receipts
type CandidateTopSectorDetails struct {
	CandidateName string `json:"cand_name"`
	Cid           string `json:"cid"`          // CRP ID
	Cycle         int    `json:"cycle,string"` // Cycle year of data being returned
	Origin        string `json:"origin"`       // Attribution to display
	Source        string `json:"source"`       // Link to CRP data
	LastUpdated   string `json:"notice"`       // Date data was retrieved from government sources (MM/DD/YYYY)
	Sectors       []Sector
}

type Sector struct {
	Name        string  `json:"sector_name"`   // CRP Sector name [Agribusiness, Communic/Electronics, Construction, Defense, Energy/Nat Resource, Finance/Insur/RealEst, Health, Lawyers & Lobbyists, Transportation, Misc Business, Labor, Ideology/Single-Issue, Other]
	Id          string  `json:"sectorid"`      // CRP's sector ID
	Total       float64 `json:"total,string"`  // Total itemized contributions attributed
	Pacs        float64 `json:"pacs,string"`   // Total contributed by PACs within sector
	Individuals float64 `json:"indivs,string"` // Total contributed by individuals within sector
}
