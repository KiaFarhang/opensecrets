package client

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/KiaFarhang/opensecrets/internal/parse"
	"github.com/KiaFarhang/opensecrets/internal/test"
	"github.com/go-playground/validator/v10"
)

const api_key string = "1"

type mockHttpClient struct {
	mockResponse http.Response
	mockError    error
}

type mockValidator struct {
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return &m.mockResponse, m.mockError
}

func (m *mockValidator) Struct(s interface{}) error {
	return nil
}

func TestGetLegislators(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := GetLegislatorsRequest{}
		_, err := client.GetLegislators(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetMemberPFDProfile(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := GetMemberPFDRequest{Year: 2020}
		_, err := client.GetMemberPFDProfile(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateSummary(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := GetCandidateSummaryRequest{Cycle: 2022}
		_, err := client.GetCandidateSummary(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateContributors(t *testing.T) {
	t.Run("Returns an error if the request passed is invaid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := GetCandidateContributorsRequest{}
		_, err := client.GetCandidateContributors(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateIndustries(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := GetCandidateIndustriesRequest{}
		_, err := client.GetCandidateIndustries(request)
		test.AssertErrorExists(err, t)
	})
}

func TestMakeGETRequest(t *testing.T) {
	t.Run("Returns an error if the HTTP call fails", func(t *testing.T) {
		mockError := errors.New("fail")
		client := openSecretsClient{client: &mockHttpClient{mockError: mockError}, validator: &mockValidator{}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		test.AssertErrorExists(err, t)
		test.AssertErrorMessage(err, "fail", t)
	})
	t.Run("Returns an error if the HTTP call is a >= 400 status code", func(t *testing.T) {
		mockResponse := buildMockResponse(400, "")
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		test.AssertErrorExists(err, t)
		wantedErrorMessage := "received 400 status code calling OpenSecrets API"
		test.AssertErrorMessage(err, wantedErrorMessage, t)
	})
	t.Run("Returns an error if the response body can't be parsed", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `BAD JSON WEEEE`)
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		test.AssertErrorExists(err, t)
		wantedErrorMessage := parse.Unable_to_parse_error_message
		test.AssertErrorMessage(err, wantedErrorMessage, t)
	})
}

func TestBuildGetLegislatorsURL(t *testing.T) {
	t.Run("Includes id passed in with request", func(t *testing.T) {
		id := "NJ"
		url := buildGetLegislatorsURL(GetLegislatorsRequest{id}, api_key)
		expectedUrl := base_url + "?method=getLegislators&output=json&apikey=" + api_key + "&id=" + id
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetMemberPFDURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := GetMemberPFDRequest{Cid: cid}
		url := buildGetMemberPFDURL(request, api_key)
		expectedUrl := base_url + "?method=memPFDProfile&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes year passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		year := 2020
		request := GetMemberPFDRequest{Cid: cid, Year: year}
		url := buildGetMemberPFDURL(request, api_key)
		expectedUrl := base_url + "?method=memPFDProfile&output=json&apikey=" + api_key + "&cid=" + cid + "&year=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateSummaryURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := GetCandidateSummaryRequest{Cid: cid}
		url := buildGetCandidateSummaryURL(request, api_key)
		expectedUrl := base_url + "?method=candSummary&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2020
		request := GetCandidateSummaryRequest{Cid: cid, Cycle: cycle}
		url := buildGetCandidateSummaryURL(request, api_key)
		expectedUrl := base_url + "?method=candSummary&output=json&apikey=" + api_key + "&cid=" + cid + "&cycle=2020"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateContributorsURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := GetCandidateContributorsRequest{Cid: cid}
		url := buildGetCandidateContributorsURL(request, api_key)
		expectedUrl := base_url + "?method=candContrib&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2022
		request := GetCandidateContributorsRequest{Cid: cid, Cycle: cycle}
		url := buildGetCandidateContributorsURL(request, api_key)
		expectedUrl := base_url + "?method=candContrib&output=json&apikey=" + api_key + "&cid=" + cid + "&cycle=2022"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateIndustriesURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := GetCandidateIndustriesRequest{Cid: cid}
		url := buildGetCandidateIndustriesURL(request, api_key)
		expectedUrl := base_url + "?method=candIndustry&output=json&apikey=" + api_key + "&cid=" + cid
		test.AssertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		request := GetCandidateIndustriesRequest{Cid: cid, Cycle: 2018}
		url := buildGetCandidateIndustriesURL(request, api_key)
		expectedUrl := base_url + "?method=candIndustry&output=json&apikey=" + api_key + "&cid=" + cid + "&cycle=2018"
		test.AssertStringMatches(url, expectedUrl, t)
	})
}

func buildMockResponse(statusCode int, jsonBody string) http.Response {
	return http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(jsonBody))}
}
