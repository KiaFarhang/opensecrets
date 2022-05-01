package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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
		request := models.LegislatorsRequest{}
		_, err := client.GetLegislators(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetMemberPFDProfile(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.MemberPFDRequest{Year: 2020}
		_, err := client.GetMemberPFDProfile(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateSummary(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.CandidateSummaryRequest{Cycle: 2022}
		_, err := client.GetCandidateSummary(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateContributors(t *testing.T) {
	t.Run("Returns an error if the request passed is invaid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.CandidateContributorsRequest{}
		_, err := client.GetCandidateContributors(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateIndustries(t *testing.T) {
	t.Run("Returns an error if the request passed is invalid", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.CandidateIndustriesRequest{}
		_, err := client.GetCandidateIndustries(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCandidateIndustryDetails(t *testing.T) {
	t.Run("Returns an error if the request doesn't have a CID", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.CandidateIndustryDetailsRequest{Ind: "K02"}
		_, err := client.GetCandidateIndustryDetails(request)
		test.AssertErrorExists(err, t)
	})
	t.Run("Returns an error if the request doesn't have an industry code", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.CandidateIndustryDetailsRequest{Cid: "N00007360"}
		_, err := client.GetCandidateIndustryDetails(request)
		test.AssertErrorExists(err, t)

	})
}

func TestGetCandidateTopSectorDetails(t *testing.T) {
	t.Run("Returns an error if the request doesn't have a CID", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.CandidateTopSectorsRequest{}
		_, err := client.GetCandidateTopSectorDetails(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetCommitteeFundraisingDetails(t *testing.T) {
	t.Run("Returns an error if the request doesn't have a committee ID", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.FundraisingByCongressionalCommitteeRequest{Industry: "ABC"}
		_, err := client.GetCommitteeFundraisingDetails(request)
		test.AssertErrorExists(err, t)
	})
	t.Run("Returns an error if the request doesn't have an industry ID", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.FundraisingByCongressionalCommitteeRequest{Committee: "HARM"}
		_, err := client.GetCommitteeFundraisingDetails(request)
		test.AssertErrorExists(err, t)
	})
}

func TestSearchForOrganization(t *testing.T) {
	t.Run("Returns an error if the request doesn't have an org name", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.OrganizationSearch{}
		_, err := client.SearchForOrganization(request)
		test.AssertErrorExists(err, t)
	})
}

func TestGetOrganizationSummary(t *testing.T) {
	t.Run("Returns an error if the request doesn't have an org ID", func(t *testing.T) {
		client := openSecretsClient{client: &mockHttpClient{}, validator: validator.New()}
		request := models.OrganizationSummaryRequest{}
		_, err := client.GetOrganizationSummary(request)
		test.AssertErrorExists(err, t)
	})
}

func TestMakeGETRequest(t *testing.T) {
	t.Run("Returns an error if the HTTP call fails", func(t *testing.T) {
		mockError := errors.New("fail")
		client := openSecretsClient{client: &mockHttpClient{mockError: mockError}, validator: &mockValidator{}}
		_, err := client.GetLegislators(models.LegislatorsRequest{})
		test.AssertErrorExists(err, t)
		test.AssertErrorMessage(err, "fail", t)
	})
	t.Run("Returns an error if the HTTP call is a >= 400 status code", func(t *testing.T) {
		mockResponse := buildMockResponse(400, "")
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(models.LegislatorsRequest{})
		test.AssertErrorExists(err, t)
		wantedErrorMessage := "received 400 status code calling OpenSecrets API"
		test.AssertErrorMessage(err, wantedErrorMessage, t)
	})
	t.Run("Returns an error if the response body can't be parsed", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `BAD JSON WEEEE`)
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		_, err := client.GetLegislators(models.LegislatorsRequest{})
		test.AssertErrorExists(err, t)
		wantedErrorMessage := parse.UnableToParseErrorMessage
		test.AssertErrorMessage(err, wantedErrorMessage, t)
	})
}

func TestMakeGetRequestWithContext(t *testing.T) {
	t.Run("returns an error if the context passed is canceled before the request completes", func(t *testing.T) {
		if testing.Short() {
			t.Skip()
		}
		client := &http.Client{}
		openSecretsClient := &openSecretsClient{client: client, validator: &mockValidator{}}
		serverClosed := make(chan bool, 1)
		handler := func(w http.ResponseWriter, r *http.Request) {
			<-serverClosed
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("hi"))
		}
		testServer := httptest.NewServer(http.HandlerFunc(handler))
		defer testServer.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		_, err := openSecretsClient.makeGETRequestWithContext(ctx, testServer.URL)

		test.AssertErrorExists(err, t)
		if !strings.Contains(err.Error(), "context deadline exceeded") {
			t.Errorf("Wanted an error containing 'context deadline exceeded' but got %s", err.Error())
		}
		serverClosed <- true
	})
}

func buildMockResponse(statusCode int, jsonBody string) http.Response {
	return http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(jsonBody))}
}
