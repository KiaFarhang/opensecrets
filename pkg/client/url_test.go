package client

import (
	"testing"

	"github.com/KiaFarhang/opensecrets/internal/test"
	"github.com/KiaFarhang/opensecrets/pkg/models"
)

const apiKey string = "1"

func TestBuildGetLegislatorsURL(t *testing.T) {
	t.Run("Includes id passed in with request", func(t *testing.T) {
		id := "NJ"
		url := buildGetLegislatorsURL(models.GetLegislatorsRequest{Id: id}, apiKey)
		expectedUrl := baseUrl + "?method=getLegislators&output=json&apikey=" + apiKey + "&id=" + id
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetMemberPFDURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetMemberPFDRequest{Cid: cid}
		url := buildGetMemberPFDURL(request, apiKey)
		expectedUrl := baseUrl + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes year passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		year := 2020
		request := models.GetMemberPFDRequest{Cid: cid, Year: year}
		url := buildGetMemberPFDURL(request, apiKey)
		expectedUrl := baseUrl + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + cid + "&year=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateSummaryURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateSummaryRequest{Cid: cid}
		url := buildGetCandidateSummaryURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2020
		request := models.GetCandidateSummaryRequest{Cid: cid, Cycle: cycle}
		url := buildGetCandidateSummaryURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + cid + "&cycle=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateContributorsURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateContributorsRequest{Cid: cid}
		url := buildGetCandidateContributorsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2022
		request := models.GetCandidateContributorsRequest{Cid: cid, Cycle: cycle}
		url := buildGetCandidateContributorsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + cid + "&cycle=2022"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateIndustriesURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateIndustriesRequest{Cid: cid}
		url := buildGetCandidateIndustriesURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateIndustriesRequest{Cid: cid, Cycle: 2018}
		url := buildGetCandidateIndustriesURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + cid + "&cycle=2018"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateIndustryDetailsURL(t *testing.T) {
	t.Run("Includes cid and industry code passed in requerst", func(t *testing.T) {
		cid := "N00007360"
		industryCode := "K02"
		request := models.GetCandidateIndustryDetailsRequest{Cid: cid, Ind: industryCode}
		url := buildGetCandidateIndustryDetailsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndByInd&output=json&apikey=" + apiKey + "&cid=" + cid + "&ind=" + industryCode
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		industryCode := "K02"
		request := models.GetCandidateIndustryDetailsRequest{Cid: cid, Ind: industryCode, Cycle: 2020}
		url := buildGetCandidateIndustryDetailsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndByInd&output=json&apikey=" + apiKey + "&cid=" + cid + "&ind=" + industryCode + "&cycle=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateTopSectorsURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateTopSectorsRequest{Cid: cid}
		url := buildGetCandidatTopSectorsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candSector&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		request := models.GetCandidateTopSectorsRequest{Cid: cid, Cycle: 2020}
		url := buildGetCandidatTopSectorsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candSector&output=json&apikey=" + apiKey + "&cid=" + cid + "&cycle=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildFundraisingByCongressionalCommitteeRequestURL(t *testing.T) {
	t.Run("Includes committee ID and industry code passed in request", func(t *testing.T) {
		committeeId := "HARM"
		industryCode := "F10"
		request := models.FundraisingByCongressionalCommitteeRequest{Committee: committeeId, Industry: industryCode}
		url := buildFundraisingByCongressionalCommitteeRequestURL(request, apiKey)
		expectedUrl := baseUrl + "?method=congCmteIndus&output=json&apikey=" + apiKey + "&cmte=" + committeeId + "&indus=" + industryCode
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes congress number passed in request if it's a non-zero value", func(t *testing.T) {
		committeeId := "HARM"
		industryCode := "F10"
		congressNumber := 116
		request := models.FundraisingByCongressionalCommitteeRequest{Committee: committeeId, Industry: industryCode, CongressNumber: congressNumber}
		url := buildFundraisingByCongressionalCommitteeRequestURL(request, apiKey)
		expectedUrl := baseUrl + "?method=congCmteIndus&output=json&apikey=" + apiKey + "&cmte=" + committeeId + "&indus=" + industryCode + "&congno=116"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}
