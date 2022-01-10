package opensecrets

type CandidateContributorSummary struct {
	CandidateName string `json:"cand_name"`
	Cid           string `json:"cid"`
	Cycle         int    `json:"cycle,string"`
	Origin        string `json:"origin"`
	Source        string `json:"source"`
	Notice        string `json:"notice"`
	Contributors  []CandidateContributor
}

type CandidateContributor struct {
	OrganizationName string  `json:"org_name"`
	Total            float64 `json:"total,string"`
	Pacs             float64 `json:"pacs,string"`
	Individuals      float64 `json:"indivs,string"`
}

// type candidateContributorSummaryWrapper struct {
// 	Attributes CandidateContributorSummary `json:"@attributes"`
// }

type candidateContributorWrapper struct {
	Attributes CandidateContributor `json:"@attributes"`
}

type contributorResponse struct {
	Wrapper []candidateContributorWrapper `json:"contributor"`
}

// type candidateContributorSummaryResponse struct {
// 	SummaryWrapper candidateContributorSummaryWrapper `json:"`
// }

type candidateContributorSummaryInnerWrapper struct {
	Attributes          CandidateContributorSummary `json:"@attributes"`
	ContributorResponse contributorResponse         `json:"contributor"`
}

type candidateContributorSummaryOutterWrapper struct {
	Contributors candidateContributorSummaryInnerWrapper `json:"contributors"`
}

type candidateContributorSummaryResponseWrapper struct {
	Response candidateContributorSummaryOutterWrapper `json:"response"`
}
