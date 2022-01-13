package client

import (
	"strconv"
	"strings"
)

const baseUrl string = "http://www.opensecrets.org/api/"

func buildGetLegislatorsURL(request GetLegislatorsRequest, apiKey string) string {
	return baseUrl + "?method=getLegislators&output=json&apikey=" + apiKey + "&id=" + request.Id
}

func buildGetMemberPFDURL(request GetMemberPFDRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Year != 0 {
		builder.WriteString("&year=")
		builder.WriteString(strconv.Itoa(request.Year))
	}

	return builder.String()
}

func buildGetCandidateSummaryURL(request GetCandidateSummaryRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateContributorsURL(request GetCandidateContributorsRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateIndustriesURL(request GetCandidateIndustriesRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(baseUrl + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}
