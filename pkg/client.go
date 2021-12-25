package opensecrets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const base_url string = "http://www.opensecrets.org/api/"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OpenSecretsClient struct {
	httpClient HttpClient
	apiKey     string
}

type GetLegislatorsRequest struct {
	id string // (Required) two-character state code or specific CID
}

func (o *OpenSecretsClient) GetLegislators(details GetLegislatorsRequest) ([]Legislator, error) {
	url := o.buildGetLegislatorsURL(details)
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

	bodyAsBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var legislators = []Legislator{}

	err = json.Unmarshal(bodyAsBytes, &legislators)

	if err != nil {
		return nil, fmt.Errorf("unable to parse response body")
	}

	return legislators, nil
}

func (o *OpenSecretsClient) buildGetLegislatorsURL(request GetLegislatorsRequest) string {
	return base_url + "?method=getLegislators&output=json&apikey=" + o.apiKey + "&id=" + request.id
}
