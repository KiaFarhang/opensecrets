package client

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/KiaFarhang/opensecrets/internal/parse"
	"github.com/KiaFarhang/opensecrets/internal/test"
	"github.com/KiaFarhang/opensecrets/pkg/models"
	"github.com/go-playground/validator/v10"
)

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
		request := models.GetLegislatorsRequest{}
		_, err := client.GetLegislators(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetMemberPFDProfile(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.GetMemberPFDRequest{Year: 2020}
		_, err := client.GetMemberPFDProfile(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateSummary(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.GetCandidateSummaryRequest{Cycle: 2022}
		_, err := client.GetCandidateSummary(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateContributors(t *testing.T) {
	t.Run("Returns an error if the request passed is invaid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.GetCandidateContributorsRequest{}
		_, err := client.GetCandidateContributors(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateIndustries(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.GetCandidateIndustriesRequest{}
		_, err := client.GetCandidateIndustries(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateIndustryDetails(t *testing.T) {
	t.Run("Returns an error if the request doesn't have a CID", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.GetCandidateIndustryDetailsRequest{Ind: "K02"}
		_, err := client.GetCandidateIndustryDetails(request)
		test.AssertErrorExists(err, t)
	})
	t.Run("Returns an error if the request doesn't have an industry code", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.GetCandidateIndustryDetailsRequest{Cid: "N00007360"}
		_, err := client.GetCandidateIndustryDetails(request)
		test.AssertErrorExists(err, t)

	})
}

func TestGetCandidateTopSectorDetails(t *testing.T) {
	t.Run("Returns an error if the request doesn't have a CID", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.GetCandidateTopSectorsRequest{}
		_, err := client.GetCandidateTopSectorDetails(request)
		test.AssertErrorExists(err, t)
	})
}

func TestMakeGETRequest(t *testing.T) {
	t.Run("Returns an error if the HTTP call fails", func(t *testing.T) {
		mockError := errors.New("fail")
		client := openSecretsClient{client: &mockHttpClient{mockError: mockError}, validator: &mockValidator{}}
		_, err := client.GetLegislators(models.GetLegislatorsRequest{})
		test.AssertErrorExists(err, t)
		test.AssertErrorMessage(err, "fail", t)
	})
	t.Run("Returns an error if the HTTP call is a >= 400 status code", func(t *testing.T) {
		mockResponse := buildMockResponse(400, "")
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(models.GetLegislatorsRequest{})
		test.AssertErrorExists(err, t)
		wantedErrorMessage := "received 400 status code calling OpenSecrets API"
		test.AssertErrorMessage(err, wantedErrorMessage, t)
	})
	t.Run("Returns an error if the response body can't be parsed", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `BAD JSON WEEEE`)
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(models.GetLegislatorsRequest{})
		test.AssertErrorExists(err, t)
		wantedErrorMessage := parse.UnableToParseErrorMessage
		test.AssertErrorMessage(err, wantedErrorMessage, t)
	})
}

func buildMockResponse(statusCode int, jsonBody string) http.Response {
	return http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(jsonBody))}
}
