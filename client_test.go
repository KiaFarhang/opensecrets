package opensecrets

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

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
		assertErrorExists(err, t)
	})
}

func TestGetMemberPFDProfile(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := GetMemberPFDRequest{Year: 2020}
		_, err := client.GetMemberPFDProfile(request)
		assertErrorExists(err, t)
	})
}

func TestGetCandidateSummary(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := GetCandidateSummaryRequest{Cycle: 2022}
		_, err := client.GetCandidateSummary(request)
		assertErrorExists(err, t)
	})
}

func TestMakeGETRequest(t *testing.T) {
	t.Run("Returns an error if the HTTP call fails", func(t *testing.T) {
		mockError := errors.New("fail")
		client := openSecretsClient{client: &mockHttpClient{mockError: mockError}, validator: &mockValidator{}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		assertErrorExists(err, t)
		assertErrorMessage(err, "fail", t)
	})
	t.Run("Returns an error if the HTTP call is a >= 400 status code", func(t *testing.T) {
		mockResponse := buildMockResponse(400, "")
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		assertErrorExists(err, t)
		wantedErrorMessage := "received 400 status code calling OpenSecrets API"
		assertErrorMessage(err, wantedErrorMessage, t)
	})
	t.Run("Returns an error if the response body can't be parsed", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `BAD JSON WEEEE`)
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		assertErrorExists(err, t)
		wantedErrorMessage := unable_to_parse_error_message
		assertErrorMessage(err, wantedErrorMessage, t)
	})
}

func TestBuildGetLegislatorsURL(t *testing.T) {
	t.Run("Includes id passed in with request", func(t *testing.T) {
		id := "NJ"
		url := buildGetLegislatorsURL(GetLegislatorsRequest{id}, api_key)
		expectedUrl := base_url + "?method=getLegislators&output=json&apikey=" + api_key + "&id=" + id
		assertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetMemberPFDURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := GetMemberPFDRequest{Cid: cid}
		url := buildGetMemberPFDURL(request, api_key)
		expectedUrl := base_url + "?method=memPFDProfile&output=json&apikey=" + api_key + "&cid=" + cid
		assertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes year passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		year := 2020
		request := GetMemberPFDRequest{Cid: cid, Year: year}
		url := buildGetMemberPFDURL(request, api_key)
		expectedUrl := base_url + "?method=memPFDProfile&output=json&apikey=" + api_key + "&cid=" + cid + "&year=2020"
		assertStringMatches(url, expectedUrl, t)
	})
}

func TestBuildGetCandidateSummaryURL(t *testing.T) {
	t.Run("Includes cid passed in request", func(t *testing.T) {
		cid := "N00007360"
		request := GetCandidateSummaryRequest{Cid: cid}
		url := buildGetCandidateSummaryURL(request, api_key)
		expectedUrl := base_url + "?method=candSummary&output=json&apikey=" + api_key + "&cid=" + cid
		assertStringMatches(url, expectedUrl, t)
	})
	t.Run("Includes cycle passed in request if it's a non-zero value", func(t *testing.T) {
		cid := "N00007360"
		cycle := 2020
		request := GetCandidateSummaryRequest{Cid: cid, Cycle: cycle}
		url := buildGetCandidateSummaryURL(request, api_key)
		expectedUrl := base_url + "?method=candSummary&output=json&apikey=" + api_key + "&cid=" + cid + "&cycle=2020"
		assertStringMatches(url, expectedUrl, t)
	})
}

func buildMockResponse(statusCode int, jsonBody string) http.Response {
	return http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(jsonBody))}
}
