package parse

import (
	"encoding/json"
	"errors"

	"github.com/KiaFarhang/opensecrets/pkg/models"
)

const UnableToParseErrorMessage string = "unable to parse OpenSecrets response body"

func ParseLegislatorsJSON(jsonBytes []byte) ([]models.Legislator, error) {

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
		return nil, errors.New(UnableToParseErrorMessage)
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
		return memberProfile, errors.New(UnableToParseErrorMessage)
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
		return models.CandidateSummary{}, errors.New(UnableToParseErrorMessage)
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
		return models.CandidateContributorSummary{}, errors.New(UnableToParseErrorMessage)
	}

	var contributors []models.CandidateContributor

	for _, contributor := range responseWrapper.Response.Contributors.Contributors {
		contributors = append(contributors, contributor.Attributes)
	}

	summary := responseWrapper.Response.Contributors.Attributes
	summary.Contributors = contributors

	return summary, nil
}

func ParseCandidateIndustriesJSON(jsonBody []byte) (models.CandidateIndustriesSummary, error) {
	type candidateIndustriesResponse struct {
		Response struct {
			Industries struct {
				Attributes models.CandidateIndustriesSummary `json:"@attributes"`
				Industry   []struct {
					Attributes models.Industry `json:"@attributes"`
				} `json:"industry"`
			} `json:"industries"`
		} `json:"response"`
	}

	var responseWrapper candidateIndustriesResponse
	err := json.Unmarshal(jsonBody, &responseWrapper)

	if err != nil {
		return models.CandidateIndustriesSummary{}, errors.New(UnableToParseErrorMessage)
	}

	summary := responseWrapper.Response.Industries.Attributes

	for _, industry := range responseWrapper.Response.Industries.Industry {
		summary.Industries = append(summary.Industries, industry.Attributes)
	}

	return summary, nil
}

func ParseCandidateIndustryDetailsJSON(jsonBody []byte) (models.CandidateIndustryDetails, error) {
	type candidateIndustryDetailsResponse struct {
		Response struct {
			Wrapper struct {
				Attributes models.CandidateIndustryDetails `json:"@attributes"`
			} `json:"candIndus"`
		} `json:"response"`
	}

	var responseWrapper candidateIndustryDetailsResponse
	err := json.Unmarshal(jsonBody, &responseWrapper)

	if err != nil {
		return models.CandidateIndustryDetails{}, errors.New(UnableToParseErrorMessage)
	}

	return responseWrapper.Response.Wrapper.Attributes, nil
}

func ParseCandidateTopSectorsJSON(jsonBody []byte) (models.CandidateTopSectorDetails, error) {
	type candidateSectorsResponse struct {
		Response struct {
			Wrapper struct {
				CandidateDetails models.CandidateTopSectorDetails `json:"@attributes"`
				SectorList       []struct {
					SectorAttributes models.Sector `json:"@attributes"`
				} `json:"sector"`
			} `json:"sectors"`
		} `json:"response"`
	}

	var responseWrapper candidateSectorsResponse
	err := json.Unmarshal(jsonBody, &responseWrapper)

	if err != nil {
		return models.CandidateTopSectorDetails{}, errors.New(UnableToParseErrorMessage)
	}

	details := responseWrapper.Response.Wrapper.CandidateDetails

	for _, sector := range responseWrapper.Response.Wrapper.SectorList {
		details.Sectors = append(details.Sectors, sector.SectorAttributes)
	}

	return details, nil
}

func ParseFundraisingByCommitteeJSON(jsonBody []byte) (models.CommitteeFundraisingDetails, error) {
	type fundraisingByCommitteeResponse struct {
		Response struct {
			Wrapper struct {
				CommitteeDetails models.CommitteeFundraisingDetails `json:"@attributes"`
				MemberList       []struct {
					Member models.CommitteeMember `json:"@attributes"`
				} `json:"member"`
			} `json:"committee"`
		} `json:"response"`
	}

	var responseWrapper fundraisingByCommitteeResponse
	err := json.Unmarshal(jsonBody, &responseWrapper)

	if err != nil {
		return models.CommitteeFundraisingDetails{}, errors.New(UnableToParseErrorMessage)
	}

	details := responseWrapper.Response.Wrapper.CommitteeDetails

	for _, member := range responseWrapper.Response.Wrapper.MemberList {
		details.Members = append(details.Members, member.Member)
	}

	return details, nil
}

func ParseOrganizationSearchJSON(jsonBody []byte) ([]models.OrganizationSearchResult, error) {
	type organizationSearchResponse struct {
		Response struct {
			Wrapper []struct {
				Attributes models.OrganizationSearchResult `json:"@attributes"`
			} `json:"organization"`
		} `json:"response"`
	}

	var responseWrapper organizationSearchResponse
	var toReturn []models.OrganizationSearchResult

	err := json.Unmarshal(jsonBody, &responseWrapper)

	if err != nil {
		return toReturn, errors.New(UnableToParseErrorMessage)
	}

	for _, result := range responseWrapper.Response.Wrapper {
		toReturn = append(toReturn, result.Attributes)
	}

	return toReturn, nil
}

func ParseOrganizationSummaryJSON(jsonBody []byte) (models.OrganizationSummary, error) {
	type organizationSummaryResponse struct {
		Response struct {
			Wrapper struct {
				Attributes models.OrganizationSummary `json:"@attributes"`
			} `json:"organization"`
		} `json:"response"`
	}

	var responseWrapper organizationSummaryResponse
	err := json.Unmarshal(jsonBody, &responseWrapper)

	if err != nil {
		return models.OrganizationSummary{}, errors.New(UnableToParseErrorMessage)
	}

	return responseWrapper.Response.Wrapper.Attributes, nil
}
