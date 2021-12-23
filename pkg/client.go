package opensecrets

import (
	"errors"
	"net/http"
)

const BASE_URL string = "http://www.opensecrets.org/api/"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OpenSecretsClient struct {
	httpClient HttpClient
	apiKey     string
}

func (o *OpenSecretsClient) GetLegislators() ([]Legislator, error) {
	url := BASE_URL + "?method=getLegislators&output=json&id=TX&apikey=" + o.apiKey
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// The API blocks requests without a user agent
	request.Header.Set("User-Agent", "Golang")

	_, err = o.httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	return nil, errors.New("it broke")
}
