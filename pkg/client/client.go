/*
Package client provides a client for the OpenSecrets REST API.
*/
package client

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/KiaFarhang/opensecrets/internal/parse"
	"github.com/KiaFarhang/opensecrets/pkg/models"
	"github.com/go-playground/validator/v10"
)

const base_url string = "http://www.opensecrets.org/api/"

/*
The OpenSecretsClient interface is responsible for communicating with the OpenSecrets REST API. The NewOpenSecretsClient
and NewOpenSecretsClientWithHttpClient functions in this package let users construct an instance of this interface.

An OpenSecretsClient is thread safe and you should use/share one throughout your application.
*/
type OpenSecretsClient interface {
	// Calls and returns the response from the getLegislators endpoint
	// https://www.opensecrets.org/api/?method=getLegislators&output=doc
	GetLegislators(request GetLegislatorsRequest) ([]models.Legislator, error)
	// Calls and returns the response from the memPFDprofile endpoint
	// https://www.opensecrets.org/api/?method=memPFDprofile&output=doc
	GetMemberPFDProfile(request GetMemberPFDRequest) (models.MemberProfile, error)
	// Calls and returns the response from the candSummary endpoint
	// https://www.opensecrets.org/api/?method=candSummary&output=doc
	GetCandidateSummary(request GetCandidateSummaryRequest) (models.CandidateSummary, error)
	// Calls and returns the response from the candContrib endpoint
	// https://www.opensecrets.org/api/?method=candContrib&output=doc
	GetCandidateContributors(request GetCandidateContributorsRequest) (models.CandidateContributorSummary, error)
}

/*
The OpenSecretsHttpClient interface lets users customize the HTTP client their OpenSecretsClient uses to communicate
with the OpenSecrets REST API. (e.g. if you have an existing HTTP client with custom logging, timeouts, etc.)

If you want to pass your own HTTP client to the OpenSecrets client, use NewOpenSecretsClientWithHttpClient. Otherwise, use
NewOpenSecretsClient and the client will use an http.Client with a 5-second timeout.
*/
type OpenSecretsHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type structValidator interface {
	Struct(s interface{}) error
}

type openSecretsClient struct {
	client    OpenSecretsHttpClient
	apiKey    string
	validator structValidator
}

type GetLegislatorsRequest struct {
	Id string `validate:"required"` // Required. Two-character specific state code, or CRP candidate ID.
}

type GetMemberPFDRequest struct {
	Cid  string `validate:"required"` // Required. CRP Candidate ID.
	Year int    // Optional. 2013, 2014, 2015 and 2016 data provided where available.
}

type GetCandidateSummaryRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID.
	Cycle int    // Optional; defaults to most recent cycle
}

type GetCandidateContributorsRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID.
	Cycle int    // Optional; defaults to most recent cycle
}

type GetCandidateIndustriesRequest struct {
	Cid   string `validate:"required"` // Required. CRP Candidate ID
	Cycle int    // Optional; defaults to most recent cycle
}

// Construct an OpenSecretsClient with the provided API key and a default http.Client (with a timeout of 5 seconds).
func NewOpenSecretsClient(apikey string) OpenSecretsClient {
	return &openSecretsClient{apiKey: apikey, client: &http.Client{Timeout: time.Second * 5}, validator: validator.New()}
}

// Construct an OpenSecretsClient with the provided API key and a custom HTTP client.
func NewOpenSecretsClientWithHttpClient(apikey string, client OpenSecretsHttpClient) OpenSecretsClient {
	return &openSecretsClient{apiKey: apikey, client: client, validator: validator.New()}
}

func (o *openSecretsClient) GetLegislators(request GetLegislatorsRequest) ([]models.Legislator, error) {

	err := o.validator.Struct(request)

	if err != nil {
		return nil, err
	}
	url := buildGetLegislatorsURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return nil, err
	}

	return parse.ParseGetLegislatorsJSON(responseBody)
}

func (o *openSecretsClient) GetMemberPFDProfile(request GetMemberPFDRequest) (models.MemberProfile, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return models.MemberProfile{}, err
	}

	url := buildGetMemberPFDURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return models.MemberProfile{}, err
	}

	return parse.ParseMemberPFDJSON(responseBody)
}

func (o *openSecretsClient) GetCandidateSummary(request GetCandidateSummaryRequest) (models.CandidateSummary, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return models.CandidateSummary{}, err
	}

	url := buildGetCandidateSummaryURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return models.CandidateSummary{}, nil
	}

	return parse.ParseCandidateSummaryJSON(responseBody)
}

func (o *openSecretsClient) GetCandidateContributors(request GetCandidateContributorsRequest) (models.CandidateContributorSummary, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return models.CandidateContributorSummary{}, err
	}

	url := buildGetCandidateContributorsURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return models.CandidateContributorSummary{}, err
	}

	return parse.ParseCandidateContributorsJSON(responseBody)
}

func (o *openSecretsClient) makeGETRequest(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// The API blocks requests without a user agent
	request.Header.Set("User-Agent", "Golang")

	response, err := o.client.Do(request)

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

func buildGetMemberPFDURL(request GetMemberPFDRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(base_url + "?method=memPFDProfile&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Year != 0 {
		builder.WriteString("&year=")
		builder.WriteString(strconv.Itoa(request.Year))
	}

	return builder.String()
}

func buildGetCandidateSummaryURL(request GetCandidateSummaryRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(base_url + "?method=candSummary&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateContributorsURL(request GetCandidateContributorsRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(base_url + "?method=candContrib&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}

func buildGetCandidateIndustriesURL(request GetCandidateIndustriesRequest, apiKey string) string {
	var builder strings.Builder
	builder.WriteString(base_url + "?method=candIndustry&output=json&apikey=" + apiKey + "&cid=" + request.Cid)

	if request.Cycle != 0 {
		builder.WriteString("&cycle=")
		builder.WriteString(strconv.Itoa(request.Cycle))
	}

	return builder.String()
}
