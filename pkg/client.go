package opensecrets

import (
	"encoding/json"
	"errors"
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
	url := buildGetLegislatorsURL(details, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return nil, err
	}

	var legislators = []Legislator{}

	err = json.Unmarshal(responseBody, &legislators)

	if err != nil {
		return nil, fmt.Errorf("unable to parse response body")
	}

	return legislators, nil
}

func (o *OpenSecretsClient) makeGETRequest(url string) ([]byte, error) {
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

	return bodyAsBytes, nil
}

func buildGetLegislatorsURL(request GetLegislatorsRequest, apiKey string) string {
	return base_url + "?method=getLegislators&output=json&apikey=" + apiKey + "&id=" + request.id
}

func parseGetLegislatorsJSON(jsonBytes []byte) ([]Legislator, error) {
	var responseWrapper = legislatorResponseWrapper{}
	err := json.Unmarshal(jsonBytes, &responseWrapper)
	if err != nil {
		return nil, errors.New("unable to parse response body")
	}
	var toReturn []Legislator
	legislatorWrappers := responseWrapper.Response.Wrapper
	for _, legislatorWrapper := range legislatorWrappers {
		toReturn = append(toReturn, legislatorWrapper.Attributes)
	}
	return toReturn, nil
}
