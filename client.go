package opensecrets

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

const base_url string = "http://www.opensecrets.org/api/"

type OpenSecretsClient interface {
	GetLegislators(request GetLegislatorsRequest) ([]Legislator, error)
	GetMemberPFDProfile(request GetMemberPFDRequest) (MemberProfile, error)
	GetCandidateSummary(request GetCandidateSummaryRequest) (CandidateSummary, error)
	GetCandidateContributors(request GetCandidateContributorsRequest) (CandidateContributorSummary, error)
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type structValidator interface {
	Struct(s interface{}) error
}

type openSecretsClient struct {
	client    httpClient
	apiKey    string
	validator structValidator
}

type GetLegislatorsRequest struct {
	Id string `validate:"required"`
}

type GetMemberPFDRequest struct {
	Cid  string `validate:"required"`
	Year int
}

type GetCandidateSummaryRequest struct {
	Cid   string `validate:"required"`
	Cycle int
}

type GetCandidateContributorsRequest struct {
	Cid   string `validate:"required"`
	Cycle int
}

func NewOpenSecretsClient(apikey string) OpenSecretsClient {
	return &openSecretsClient{apiKey: apikey, client: &http.Client{Timeout: time.Second * 5}, validator: validator.New()}
}

func NewOpenSecretsClientWithHttpClient(apikey string, client httpClient) OpenSecretsClient {
	return &openSecretsClient{apiKey: apikey, client: client, validator: validator.New()}
}

func (o *openSecretsClient) GetLegislators(request GetLegislatorsRequest) ([]Legislator, error) {

	err := o.validator.Struct(request)

	if err != nil {
		return nil, err
	}
	url := buildGetLegislatorsURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return nil, err
	}

	return parseGetLegislatorsJSON(responseBody)
}

func (o *openSecretsClient) GetMemberPFDProfile(request GetMemberPFDRequest) (MemberProfile, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return MemberProfile{}, err
	}

	url := buildGetMemberPFDURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return MemberProfile{}, err
	}

	return parseMemberPFDJSON(responseBody)
}

func (o *openSecretsClient) GetCandidateSummary(request GetCandidateSummaryRequest) (CandidateSummary, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return CandidateSummary{}, err
	}

	url := buildGetCandidateSummaryURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return CandidateSummary{}, nil
	}

	return parseCandidateSummaryJSON(responseBody)
}

func (o *openSecretsClient) GetCandidateContributors(request GetCandidateContributorsRequest) (CandidateContributorSummary, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return CandidateContributorSummary{}, err
	}

	url := buildGetCandidateContributorsURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return CandidateContributorSummary{}, err
	}

	return parseCandidateContributorsJSON(responseBody)
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
