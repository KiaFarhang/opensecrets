package parse

import (
	"encoding/json"
	"errors"

	"github.com/KiaFarhang/opensecrets/pkg/models"
)

const Unable_to_parse_error_message string = "unable to parse OpenSecrets response body"

func ParseGetLegislatorsJSON(jsonBytes []byte) ([]models.Legislator, error) {

	type legislatorResponse struct {
		Response struct {
			Legislator []struct {
				Attributes models.Legislator `json:"@attributes"`
			} `json:"legislator"`
		} `json:"response"`
	}

	var responseWrapper = legislatorResponse{}
	err := json.Unmarshal(jsonBytes, &responseWrapper)
	if err != nil {
		return nil, errors.New(Unable_to_parse_error_message)
	}

	var toReturn []models.Legislator
	for _, legislatorWrapper := range responseWrapper.Response.Legislator {
		toReturn = append(toReturn, legislatorWrapper.Attributes)
	}

	return toReturn, nil
}

func ParseMemberPFDJSON(jsonBtyes []byte) (models.MemberProfile, error) {

	type memberPFDResponse struct {
		Response struct {
			Wrapper struct {
				Profile      models.MemberProfile `json:"@attributes"`
				AssetWrapper struct {
					Assets []struct {
						Asset models.Asset `json:"@attributes"`
					} `json:"asset"`
				} `json:"assets"`
				TransactionWrapper struct {
					Transactions []struct {
						Transaction models.Transaction `json:"@attributes"`
					} `json:"transaction"`
				} `json:"transactions"`
				PositionWrapper struct {
					Positions []struct {
						Position models.Position `json:"@attributes"`
					} `json:"position"`
				} `json:"positions"`
			} `json:"member_profile"`
		} `json:"response"`
	}

	var memberProfile models.MemberProfile
	var responseWrapper = memberPFDResponse{}
	err := json.Unmarshal(jsonBtyes, &responseWrapper)
	if err != nil {
		return memberProfile, errors.New(Unable_to_parse_error_message)
	}

	memberProfile = responseWrapper.Response.Wrapper.Profile

	var memberAssets []models.Asset
	assetWrappers := responseWrapper.Response.Wrapper.AssetWrapper.Assets
	for _, assetWrapper := range assetWrappers {
		memberAssets = append(memberAssets, assetWrapper.Asset)
	}
	memberProfile.Assets = memberAssets

	var memberTransactions []models.Transaction
	transactionWrappers := responseWrapper.Response.Wrapper.TransactionWrapper.Transactions
	for _, transactionWrapper := range transactionWrappers {
		memberTransactions = append(memberTransactions, transactionWrapper.Transaction)
	}
	memberProfile.Transactions = memberTransactions

	var memberPositions []models.Position
	positionWrappers := responseWrapper.Response.Wrapper.PositionWrapper.Positions
	for _, positionWrapper := range positionWrappers {
		memberPositions = append(memberPositions, positionWrapper.Position)
	}
	memberProfile.Positions = memberPositions

	return memberProfile, nil
}

func ParseCandidateSummaryJSON(jsonBytes []byte) (models.CandidateSummary, error) {
	type candidateSummaryResponse struct {
		Response struct {
			Summary struct {
				Attributes models.CandidateSummary `json:"@attributes"`
			} `json:"summary"`
		} `json:"response"`
	}

	var responseWrapper candidateSummaryResponse
	err := json.Unmarshal(jsonBytes, &responseWrapper)
	if err != nil {
		return models.CandidateSummary{}, errors.New(Unable_to_parse_error_message)
	}
	return responseWrapper.Response.Summary.Attributes, nil
}

func ParseCandidateContributorsJSON(jsonBytes []byte) (models.CandidateContributorSummary, error) {

	type candidateContributorResponse struct {
		Response struct {
			Contributors struct {
				Attributes   models.CandidateContributorSummary `json:"@attributes"`
				Contributors []struct {
					Attributes models.CandidateContributor `json:"@attributes"`
				} `json:"contributor"`
			} `json:"contributors"`
		} `json:"response"`
	}

	var responseWrapper candidateContributorResponse
	err := json.Unmarshal(jsonBytes, &responseWrapper)
	if err != nil {
		return models.CandidateContributorSummary{}, errors.New(Unable_to_parse_error_message)
	}

	var contributors []models.CandidateContributor

	for _, contributor := range responseWrapper.Response.Contributors.Contributors {
		contributors = append(contributors, contributor.Attributes)
	}

	summary := responseWrapper.Response.Contributors.Attributes
	summary.Contributors = contributors

	return summary, nil
}
