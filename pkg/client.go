package opensecrets

import (
	"errors"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OpenSecretsClient struct {
	httpClient HttpClient
}

func (o *OpenSecretsClient) GetLegislators() ([]Legislator, error) {
	return nil, errors.New("it broke")
}
