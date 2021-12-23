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
		result := client.GetLegislators()
		if result != "foo" {
			t.Errorf("Wanted 'foo' got %s", result)
		}
	})
}
