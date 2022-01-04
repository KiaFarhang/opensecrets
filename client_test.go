package opensecrets

import (
	"errors"
	"io"
	"net/http"
	"reflect"
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
	t.Run("Returns a slice of Legislators", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `{"response": {"legislator": [{"@attributes": {"first_elected": "2000"}}, {"@attributes": {"first_elected": "2005"}}]}}`)
		client := openSecretsClient{client: &mockHttpClient{mockResponse: mockResponse}, validator: &mockValidator{}}
		legislators, err := client.GetLegislators(GetLegislatorsRequest{})
		if err != nil {
			t.Fatalf("Expected no error but got one with message %s", err.Error())
		}
		if len(legislators) != 2 {
			t.Fatalf("Expected 2 legislators but got %d", len(legislators))
		}
		expectedLegislators := []Legislator{
			{FirstElected: 2000},
			{FirstElected: 2005},
		}

		if !reflect.DeepEqual(legislators, expectedLegislators) {
			t.Fatalf("Got %v want %v", legislators, expectedLegislators)
		}
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
		wantedErrorMessage := "unable to parse response body"
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

func TestParseGetLegislatorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json := []byte(`{"response": {"legislator": [{"@attributes": {"first_elected": "2000"}}]}}`)
		legislators, err := parseGetLegislatorsJSON(json)
		if err != nil {
			t.Fatalf("Expected no error but got one with message %s", err.Error())
		}
		expectedLegislators := []Legislator{
			{FirstElected: 2000},
		}

		if !reflect.DeepEqual(legislators, expectedLegislators) {
			t.Fatalf("Got %v want %v", legislators, expectedLegislators)
		}
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := parseGetLegislatorsJSON(json)
		wantedErrorMessage := "unable to parse response body"
		assertErrorMessage(err, wantedErrorMessage, t)
	})
}

func assertErrorExists(err error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Fatalf("Wanted error but got nil")
	}
}

func assertErrorMessage(err error, expectedMessage string, t *testing.T) {
	t.Helper()
	if err.Error() != expectedMessage {
		t.Fatalf("Wanted error message %s but got %s", expectedMessage, err.Error())
	}

}

func assertStringMatches(got, wanted string, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Fatalf("Wanted string %s got string %s", wanted, got)
	}
}

func buildMockResponse(statusCode int, jsonBody string) http.Response {
	return http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(jsonBody))}
}
