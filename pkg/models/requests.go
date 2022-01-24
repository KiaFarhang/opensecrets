package models

type LegislatorsRequest struct {
	Id string `validate:"required"` // Required. Two-character specific state code, or CRP candidate ID.
}

type MemberPFDRequest struct {
	Cid  string `validate:"required"` // Required. CRP Candidate ID.
	Year int    // Optional. 2013, 2014, 2015 and 2016 data provided where available.
}

type CandidateSummaryRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID.
	Cycle int    // Optional; defaults to most recent cycle
}

type CandidateContributorsRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID.
	Cycle int    // Optional; defaults to most recent cycle
}

type CandidateIndustriesRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID
	Cycle int    // Optional; defaults to most recent cycle
}

type CandidateIndustryDetailsRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID
	Ind   string `validate:"required"` // Required. A 3-character industry code
	Cycle int    // Optional; defaults to most recent cycle
}

type CandidateTopSectorsRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID
	Cycle int    // Optional; defaults to most recent cycle
}

type FundraisingByCongressionalCommitteeRequest struct {
	Committee      string `validate:"required"` // Required. Committee ID in CQ format
	Industry       string `validate:"required"` // Required. Industry code
	CongressNumber int    // Optional, defaults to most recent Congress
}
