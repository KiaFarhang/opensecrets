package client

import (
	"strconv"
	"strings"

	"github.com/KiaFarhang/opensecrets/pkg/models"
)

const baseUrl string = "http://www.opensecrets.org/api/"

func buildLegislatorsURL(request models.LegislatorsRequest, apiKey string) string {
	return baseUrl + "?method=getLegislators&output=json&apikey=" + apiKey + "&id=" + request.Id
}

func buildMemberPFDURL(request models.MemberPFDRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Year != 0 {
		builder.WriteString("&year=")
		builder.WriteString(strconv.Itoa(request.Year))
	}

	return builder.String()
}

func buildCandidateSummaryURL(request models.CandidateSummaryRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildCandidateContributorsURL(request models.CandidateContributorsRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateIndustriesURL(request models.CandidateIndustriesRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildCandidateIndustryDetailsURL(request models.CandidateIndustryDetailsRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candIndByInd&output=json&apikey=" + apiKey + "&cid=" + request.Cid + "&ind=" + request.Ind)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildCandidateTopSectorsURL(request models.CandidateTopSectorsRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candSector&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildFundraisingByCongressionalCommitteeRequestURL(request models.FundraisingByCongressionalCommitteeRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=congCmteIndus&output=json&apikey=" + apiKey + "&cmte=" + request.Committee + "&indus=" + request.Industry)

	if request.CongressNumber != 0 {
		builder.WriteString("&congno=")
		builder.WriteString(strconv.Itoa(request.CongressNumber))
	}

	return builder.String()
}

func buildOrganizationSearchURL(request models.OrganizationSearch, apiKey string) string {
	return baseUrl + "?method=getOrgs&output=json&apikey=" + apiKey + "&org=" + request.Name
}

func buildOrganizationSummaryURL(request models.OrganizationSummaryRequest, apiKey string) string {
	return baseUrl + "?method=orgSummary&output=json&apikey=" + apiKey + "&id=" + request.Id
}

func buildIndependentExpendituresURL(apiKey string) string {
	return baseUrl + "?method=independentExpend&output=json&apikey=" + apiKey
}
