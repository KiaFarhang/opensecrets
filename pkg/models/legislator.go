package models

// A current member of Congress.
type Legislator struct {
	FirstElected   int    `json:"first_elected,string"`
	Cid            string `json:"cid"` // CRP ID for the legislator
	FirstLast      string `json:"firstlast"`
	LastName       string `json:"lastname"`
	Party          string `json:"party"`
	Office         string `json:"office"`
	Gender         string `json:"gender"`           // M or F
	ExitCode       int    `json:"exit_code,string"` // Assigned by CRP, see OpenData user's guide for details
	Comments       string `json:"comments"`         // Generally expounds on exit code
	Phone          string `json:"phone"`
	Fax            string `json:"fax"`
	Website        string `json:"website"`
	Webform        string `json:"webform"`
	CongressOffice string `json:"congress_office"`
	BioguideId     string `json:"bioguide_id"`  // ID of a member from the Congressional BioGuide
	VoteSmartId    string `json:"votesmart_id"` // VoteSmart ID of member
	FECCandId      string `json:"feccandid"`    // ID of member assigned by Federal Election Commission
	TwitterId      string `json:"twitter_id"`
	YouTubeURL     string `json:"youtube_url"`
	FacebookId     string `json:"facebook_id"`
	Birthdate      string `json:"birthdate"` // YYYY-MM-DD
}
