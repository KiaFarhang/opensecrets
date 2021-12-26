package opensecrets

type Legislator struct {
	FirstElected int `json:"first_elected,string"`
	// Cid            string `json:"cid"`
	// FirstLast      string `json:"firstlast"`
	// LastName       string `json:"lastname"`
	// Party          string `json:"party"`
	// Office         string `json:"office"`
	// Gender         string `json:"gender"`
	// ExitCode       int    `json:"exit_code,string"`
	// Comments       string `json:"comments"`
	// Phone          string `json:"phone"`
	// Fax            string `json:"fax"`
	// Website        string `json:"website"`
	// Webform        string `json:"webform"`
	// CongressOffice string `json:"congress_office"`
	// BioguideId     string `json:"bioguide_id"`
	// VoteSmartId    int    `json:"votesmart_id,string"`
	// FECCandId      string `json:"feccandid"`
	// TwitterId      string `json:"twitter_id"`
	// YouTubeURL     string `json:"youtube_url"`
	// FacebookId     int    `json:"facebook_id,string"`
	// TODO: Could make this a better type
	// https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
	// Birthdate string `json:"birthdate"`
}

type legislatorWrapper struct {
	Attributes Legislator `json:"@attributes"`
}

type legislatorResponse struct {
	Wrapper []legislatorWrapper `json:"legislator"`
}

type legislatorResponseWrapper struct {
	Response legislatorResponse `json:"response"`
}
