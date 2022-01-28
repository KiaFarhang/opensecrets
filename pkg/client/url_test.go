package client

import (
	"testing"

	"github.com/KiaFarhang/opensecrets/internal/test"
	"github.com/KiaFarhang/opensecrets/pkg/models"
)

const apiKey string = "1"

func TestBuildLegislatorsURL(t *testing.T) {
	t.Run("Includes id passed in with request", func(t *testing.T) {
		id := "NJ"
		url := buildLegislatorsURL(models.LegislatorsRequest{Id: id}, apiKey)
		expectedUrl := baseUrl + "?method=getLegislators&output=json&apikey=" + apiKey + "&id=" + id
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildMemberPFDURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.MemberPFDRequest{Cid: cid}
		url := buildMemberPFDURL(request, apiKey)
		expectedUrl := baseUrl + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes year passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		year := 2020
		request := models.MemberPFDRequest{Cid: cid, Year: year}
		url := buildMemberPFDURL(request, apiKey)
		expectedUrl := baseUrl + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + cid + "&year=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildCandidateSummaryURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.CandidateSummaryRequest{Cid: cid}
		url := buildCandidateSummaryURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2020
		request := models.CandidateSummaryRequest{Cid: cid, Cycle: cycle}
		url := buildCandidateSummaryURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + cid + "&cycle=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildCandidateContributorsURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.CandidateContributorsRequest{Cid: cid}
		url := buildCandidateContributorsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2022
		request := models.CandidateContributorsRequest{Cid: cid, Cycle: cycle}
		url := buildCandidateContributorsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + cid + "&cycle=2022"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildCandidateIndustriesURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.CandidateIndustriesRequest{Cid: cid}
		url := buildGetCandidateIndustriesURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		request := models.CandidateIndustriesRequest{Cid: cid, Cycle: 2018}
		url := buildGetCandidateIndustriesURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + cid + "&cycle=2018"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildCandidateIndustryDetailsURL(t *testing.T) {
	t.Run("Includes cid and industry code passed in requerst", func(t *testing.T) {
		cid := "N00007360"
		industryCode := "K02"
		request := models.CandidateIndustryDetailsRequest{Cid: cid, Ind: industryCode}
		url := buildCandidateIndustryDetailsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndByInd&output=json&apikey=" + apiKey + "&cid=" + cid + "&ind=" + industryCode
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		industryCode := "K02"
		request := models.CandidateIndustryDetailsRequest{Cid: cid, Ind: industryCode, Cycle: 2020}
		url := buildCandidateIndustryDetailsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candIndByInd&output=json&apikey=" + apiKey + "&cid=" + cid + "&ind=" + industryCode + "&cycle=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildCandidateTopSectorsURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := models.CandidateTopSectorsRequest{Cid: cid}
		url := buildCandidateTopSectorsURL(request, apiKey)
		expectedUrl := baseUrl + "?method=candSector&output=json&apikey=" + apiKey + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		request := models.CandidateTopSectorsRequest{Cid: cid, Cycle: 2020}
		url := buildCandidateTopSectorsURL(request, apiKey)
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

func TestBuildOrganizationSearchURL(t *testing.T) {
	t.Run("Includes organization query passed in request", func(t *testing.T) {
		org := "Foo"
		request := models.OrganizationSearch{Name: org}
		url := buildOrganizationSearchURL(request, apiKey)
		expectedUrl := baseUrl + "?method=getOrgs&output=json&apikey=" + apiKey + "&org=" + org
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildOrganizationSummaryURL(t *testing.T) {
	t.Run("Includes org ID passed in request", func(t *testing.T) {
		id := "123"
		request := models.OrganizationSummaryRequest{Id: id}
		url := buildOrganizationSummaryURL(request, apiKey)
		expectedUrl := baseUrl + "?method=orgSummary&output=json&apikey=" + apiKey + "&id=" + id
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildIndependentExpendituresURL(t *testing.T) {
	t.Run("Returns the expected URL", func(t *testing.T) {
		url := buildIndependentExpendituresURL(apiKey)
		expectedUrl := baseUrl + "?method=independentExpend&output=json&apikey=" + apiKey
		test.AssertStringMatches(url, expectedUrl, t)
	})
}
