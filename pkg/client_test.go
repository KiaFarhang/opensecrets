package opensecrets

import (
	"errors"
	"io"
	"net/http"
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
		_, err := client.GetLegislators()
		assertErrorExists(err, t)
		if err.Error() != "fail" {
			t.Fatalf("Wanted error string %s but got %s", mockError.Error(), err.Error())
		}
	})
	t.Run("Returns an error if the HTTP call is a non-200 status code", func(t *testing.T) {
		mockResponse := buildMockResponse(400, "")
		client := OpenSecretsClient{httpClient: &mockHttpClient{mockResponse: mockResponse}}
		_, err := client.GetLegislators()
		assertErrorExists(err, t)
		wantedErrorString := "received 400 status code calling OpenSecrets API"
		if err.Error() != wantedErrorString {
			t.Fatalf("Wanted error string %s but got %s", wantedErrorString, err.Error())
		}
	})
	t.Run("Returns an error if the response body can't be parsed", func(t *testing.T) {
		mockResponse := buildMockResponse(200, `BAD JSON WEEEE`)
		client := OpenSecretsClient{httpClient: &mockHttpClient{mockResponse: mockResponse}}
		_, err := client.GetLegislators()
		assertErrorExists(err, t)
		wantedErrorMessage := "unable to parse response body"
		if err.Error() != wantedErrorMessage {
			t.Fatalf("Wanted error message %s but got %s", wantedErrorMessage, err.Error())
		}
	})

}

func assertErrorExists(err error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Fatalf("Wanted error but got nil")
	}
}

func buildMockResponse(statusCode int, jsonBody string) http.Response {
	return http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(jsonBody))}
}
