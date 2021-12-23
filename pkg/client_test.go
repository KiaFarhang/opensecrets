package opensecrets

import (
	"errors"
	"net/http"
	"testing"
)

type mockHttpClient struct {
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("broken!")
}

func TestGetLegislators(t *testing.T) {
	t.Run("Dummy test", func(t *testing.T) {
		client := OpenSecretsClient{httpClient: &mockHttpClient{}}
		_, err := client.GetLegislators()
		if err == nil {
			t.Error("Wanted non-nil error but got nil")
		}
	})
}
