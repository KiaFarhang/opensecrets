package client

import (
	"testing"

	"github.com/KiaFarhang/opensecrets/internal/test"
	"github.com/KiaFarhang/opensecrets/pkg/models"
)

func TestBuildGetLegislatorsURL(t *testing.T) {
	t.Run("Includes id passed in with request", func(t *testing.T) {
		id := "NJ"
		url := buildGetLegislatorsURL(models.GetLegislatorsRequest{id}, api_key)
		expectedUrl := baseUrl + "?method=getLegislators&output=json&apikey=" + api_key + "&id=" + id
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetMemberPFDURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetMemberPFDRequest{Cid: cid}
		url := buildGetMemberPFDURL(request, api_key)
		expectedUrl := baseUrl + "?method=memPFDProfile&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes year passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		year := 2020
		request := models.GetMemberPFDRequest{Cid: cid, Year: year}
		url := buildGetMemberPFDURL(request, api_key)
		expectedUrl := baseUrl + "?method=memPFDProfile&output=json&apikey=" + api_key + "&cid=" + cid + "&year=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateSummaryURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateSummaryRequest{Cid: cid}
		url := buildGetCandidateSummaryURL(request, api_key)
		expectedUrl := baseUrl + "?method=candSummary&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2020
		request := models.GetCandidateSummaryRequest{Cid: cid, Cycle: cycle}
		url := buildGetCandidateSummaryURL(request, api_key)
		expectedUrl := baseUrl + "?method=candSummary&output=json&apikey=" + api_key + "&cid=" + cid + "&cycle=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateContributorsURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateContributorsRequest{Cid: cid}
		url := buildGetCandidateContributorsURL(request, api_key)
		expectedUrl := baseUrl + "?method=candContrib&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2022
		request := models.GetCandidateContributorsRequest{Cid: cid, Cycle: cycle}
		url := buildGetCandidateContributorsURL(request, api_key)
		expectedUrl := baseUrl + "?method=candContrib&output=json&apikey=" + api_key + "&cid=" + cid + "&cycle=2022"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateIndustriesURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateIndustriesRequest{Cid: cid}
		url := buildGetCandidateIndustriesURL(request, api_key)
		expectedUrl := baseUrl + "?method=candIndustry&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateIndustriesRequest{Cid: cid, Cycle: 2018}
		url := buildGetCandidateIndustriesURL(request, api_key)
		expectedUrl := baseUrl + "?method=candIndustry&output=json&apikey=" + api_key + "&cid=" + cid + "&cycle=2018"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}
