package opensecrets

import (
	"encoding/json"
	"errors"
)

const unable_to_parse_error_message string = "unable to parse OpenSecrets response body"

func parseGetLegislatorsJSON(jsonBytes []byte) ([]Legislator, error) {
	var responseWrapper = legislatorResponseWrapper{}
	err := json.Unmarshal(jsonBytes, &responseWrapper)
	if err != nil {
		return nil, errors.New(unable_to_parse_error_message)
	}
	var toReturn []Legislator
	legislatorWrappers := responseWrapper.Response.Wrapper
	for _, legislatorWrapper := range legislatorWrappers {
		toReturn = append(toReturn, legislatorWrapper.Attributes)
	}
	return toReturn, nil
}

func parseMemberPFDJSON(jsonBtyes []byte) (MemberProfile, error) {
	var memberProfile MemberProfile
	var responseWrapper = memberPFDResponseWrapper{}
	err := json.Unmarshal(jsonBtyes, &responseWrapper)
	if err != nil {
		return memberProfile, errors.New(unable_to_parse_error_message)
	}

	memberProfile = responseWrapper.Response.Profile.Attributes

	var memberAssets []Asset
	assetWrappers := responseWrapper.Response.Profile.Assets.Wrapper
	for _, assetWrapper := range assetWrappers {
		memberAssets = append(memberAssets, assetWrapper.Attributes)
	}
	memberProfile.Assets = memberAssets

	var memberTransactions []Transaction
	transactionWrappers := responseWrapper.Response.Profile.Transactions.Wrapper
	for _, transactionWrapper := range transactionWrappers {
		memberTransactions = append(memberTransactions, transactionWrapper.Attributes)
	}
	memberProfile.Transactions = memberTransactions

	var memberPositions []Position
	positionWrappers := responseWrapper.Response.Profile.Positions.Wrapper
	for _, positionWrapper := range positionWrappers {
		memberPositions = append(memberPositions, positionWrapper.Attributes)
	}
	memberProfile.Positions = memberPositions

	return memberProfile, nil
}

func parseCandidateSummaryJSON(jsonBytes []byte) (CandidateSummary, error) {
	var responseWrapper candidateSummaryResponseWrapper
	err := json.Unmarshal(jsonBytes, &responseWrapper)
	if err != nil {
		return CandidateSummary{}, errors.New(unable_to_parse_error_message)
	}
	return responseWrapper.Response.Wrapper.Attributes, nil
}

func parseCandidateContributorsJSON(jsonBytes []byte) (CandidateContributorSummary, error) {

	type CandidateContributorResponse struct {
		Response struct {
			Contributors struct {
				Attributes   CandidateContributorSummary `json:"@attributes"`
				Contributors []struct {
					Attributes CandidateContributor `json:"@attributes"`
				} `json:"contributor"`
			} `json:"contributors"`
		} `json:"response"`
	}

	var responseWrapper CandidateContributorResponse
	err := json.Unmarshal(jsonBytes, &responseWrapper)
	if err != nil {
		return CandidateContributorSummary{}, errors.New(unable_to_parse_error_message)
	}

	var contributors []CandidateContributor

	for _, contributor := range responseWrapper.Response.Contributors.Contributors {
		contributors = append(contributors, contributor.Attributes)
	}

	summary := responseWrapper.Response.Contributors.Attributes
	summary.Contributors = contributors

	return summary, nil
}
