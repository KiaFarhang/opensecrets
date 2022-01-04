package opensecrets

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

const base_url string = "http://www.opensecrets.org/api/"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type StructValidator interface {
	Struct(s interface{}) error
}

type openSecretsClient struct {
	httpClient HttpClient
	apiKey     string
	validator  StructValidator
}

type GetLegislatorsRequest struct {
	Id string `validate:"required"`
}

func NewOpenSecretsClient(apikey string) openSecretsClient {
	return openSecretsClient{apiKey: apikey, httpClient: &http.Client{Timeout: time.Second * 5}, validator: validator.New()}
}

func NewOpenSecretsClientWithHttpClient(apikey string, httpClient HttpClient) openSecretsClient {
	return openSecretsClient{apiKey: apikey, httpClient: httpClient, validator: validator.New()}
}

func (o *openSecretsClient) GetLegislators(details GetLegislatorsRequest) ([]Legislator, error) {

	err := o.validator.Struct(details)

	if err != nil {
		return nil, err
	}
	url := buildGetLegislatorsURL(details, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return nil, err
	}

	return parseGetLegislatorsJSON(responseBody)
}

func (o *openSecretsClient) makeGETRequest(url string) ([]byte, error) {
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
	return base_url + "?method=getLegislators&output=json&apikey=" + apiKey + "&id=" + request.Id
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
