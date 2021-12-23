package opensecrets

import "net/http"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OpenSecretsClient struct {
	httpClient HttpClient
}

func (o *OpenSecretsClient) GetLegislators() string {
	return "foo"
}
