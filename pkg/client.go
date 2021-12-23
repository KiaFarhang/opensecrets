package opensecrets

import (
	"errors"
	"fmt"
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

	response, err := o.httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	statusCode := response.StatusCode

	if statusCode >= 400 {
		return nil, fmt.Errorf("received %d status code calling OpenSecrets API", statusCode)
	}

	return nil, errors.New("it broke")
}
