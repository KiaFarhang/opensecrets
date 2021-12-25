package opensecrets

type Legislator struct {
	FirstElected int `json:"first_elected,string"`
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
