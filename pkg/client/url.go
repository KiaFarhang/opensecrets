package client

import (
	"strconv"
	"strings"

	"github.com/KiaFarhang/opensecrets/pkg/models"
)

const baseUrl string = "http://www.opensecrets.org/api/"

func buildGetLegislatorsURL(request models.GetLegislatorsRequest, apiKey string) string {
	return baseUrl + "?method=getLegislators&output=json&apikey=" + apiKey + "&id=" + request.Id
}

func buildGetMemberPFDURL(request models.GetMemberPFDRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Year != 0 {
		builder.WriteString("&year=")
		builder.WriteString(strconv.Itoa(request.Year))
	}

	return builder.String()
}

func buildGetCandidateSummaryURL(request models.GetCandidateSummaryRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateContributorsURL(request models.GetCandidateContributorsRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateIndustriesURL(request models.GetCandidateIndustriesRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateIndustryDetailsURL(request models.GetCandidateIndustryDetailsRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candIndByInd&output=json&apikey=" + apiKey + "&cid=" + request.Cid + "&ind=" + request.Ind)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}