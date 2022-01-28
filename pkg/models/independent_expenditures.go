package models

// An independent expenditure transaction
type IndependentExpenditure struct {
	CommitteeId     string  `json:"cmteid"` // ID of committee
	CommitteeName   string  `json:"pacshort"`
	SupportOrOppose string  `json:"suppopp"`       // supports (FOR:)/opposes (AGAINST:)
	CandidateName   string  `json:"candname"`      // candidate targeted
	District        string  `json:"district"`      // four-character abbreviation of district candidate is running for (e.g. NYS1)
	Amount          float64 `json:"amount,string"` // amount spent
	Note            string  `json:"note"`
	Party           string  `json:"party"` // R, D, 3, L, U (for Dem, Repub, third party, Libertarian, unknown)
	Payee           string  `json:"payee"`
	Date            string  `json:"date"`   // date of expenditure (YYYY-MM-DD HH:mm:ss.ff)
	Origin          string  `json:"origin"` //  required attribution to display
	Source          string  `json:"source"` // link to CRP web site
}
