package opensecrets

import (
	"errors"
	"net/http"
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
		if err == nil {
			t.Fatalf("Wanted error but got nil")
		}
		if err.Error() != "fail" {
			t.Fatalf("Wanted error string %s but got %s", mockError.Error(), err.Error())
		}
	})
	t.Run("Returns an error if the HTTP call is a non-200 status code", func(t *testing.T) {
		mockResponse := buildMockResponse(400)
		client := OpenSecretsClient{httpClient: &mockHttpClient{mockResponse: mockResponse}}
		_, err := client.GetLegislators()
		if err == nil {
			t.Fatalf("Wanted error but got nil")
		}
		wantedErrorString := "received 400 status code calling OpenSecrets API"
		if err.Error() != wantedErrorString {
			t.Fatalf("Wanted error string %s but got %s", wantedErrorString, err.Error())
		}
	})
}

func buildMockResponse(statusCode int) http.Response {
	return http.Response{StatusCode: statusCode}
}
