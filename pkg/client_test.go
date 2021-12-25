package opensecrets

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type mockHttpClient struct {
	mockResponse http.Response
	mockError    error
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return &m.mockResponse, m.mockError
}

func TestGetLegislators(t *testing.T) {
	t.Run("Returns an error if the HTTP call fails", func(t *testing.T) {
		mockError := errors.New("fail")
		client := OpenSecretsClient{httpClient: &mockHttpClient{mockError: mockError}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		assertErrorExists(err, t)
		assertErrorMessage(err, "fail", t)
	})
	t.Run("Returns an error if the HTTP call is a non-200 status code", func(t *testing.T) {
		mockResponse := buildMockResponse(400, "")
		client := OpenSecretsClient{httpClient: &mockHttpClient{mockResponse: mockResponse}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		assertErrorExists(err, t)
		wantedErrorMessage := "received 400 status code calling OpenSecrets API"
		assertErrorMessage(err, wantedErrorMessage, t)
	})
	t.Run("Returns an error if the response body can't be parsed", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `BAD JSON WEEEE`)
		client := OpenSecretsClient{httpClient: &mockHttpClient{mockResponse: mockResponse}}
		_, err := client.GetLegislators(GetLegislatorsRequest{})
		assertErrorExists(err, t)
		wantedErrorMessage := "unable to parse response body"
		assertErrorMessage(err, wantedErrorMessage, t)
	})
	t.Run("Returns a slice of Legislators", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `[{"first_elected": "2000"}, {"first_elected": "2005"}]`)
		client := OpenSecretsClient{httpClient: &mockHttpClient{mockResponse: mockResponse}}
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
func buildMockResponse(statusCode int, jsonBody string) http.Response {
	return http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(jsonBody))}
}
