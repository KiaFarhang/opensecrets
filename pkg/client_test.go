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
}

func buildMockResponse() http.Response {
	return http.Response{}
}
