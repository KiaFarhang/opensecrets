/*
Package client provides a client for the OpenSecrets REST API.
*/
package client

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KiaFarhang/opensecrets/internal/parse"
	"github.com/KiaFarhang/opensecrets/pkg/models"
	"github.com/go-playground/validator/v10"
)

/*
The OpenSecretsClient interface is responsible for communicating with the OpenSecrets REST API. The NewOpenSecretsClient
and NewOpenSecretsClientWithHttpClient functions in this package let users construct an instance of this interface.

An OpenSecretsClient is thread safe and you should use/share one throughout your application.
*/
type OpenSecretsClient interface {
	// Provides a list of Congressional legislators for a specified subset (state or specific CID)
	// https://www.opensecrets.org/api/?method=getLegislators&output=doc
	GetLegislators(request models.GetLegislatorsRequest) ([]models.Legislator, error)
	// Returns data on the personal finances of a member of Congress, as well as judicial + executive branches
	// https://www.opensecrets.org/api/?method=memPFDprofile&output=doc
	GetMemberPFDProfile(request models.GetMemberPFDRequest) (models.MemberProfile, error)
	// Provides summary fundraising information for a politician
	// https://www.opensecrets.org/api/?method=candSummary&output=doc
	GetCandidateSummary(request models.GetCandidateSummaryRequest) (models.CandidateSummary, error)
	// Returns top contributors to a candidate for/sitting member of Congress
	// https://www.opensecrets.org/api/?method=candContrib&output=doc
	GetCandidateContributors(request models.GetCandidateContributorsRequest) (models.CandidateContributorSummary, error)
	// Provides the top 10 industries contributing to a candidate
	// https://www.opensecrets.org/api/?method=candIndustry&output=doc
	GetCandidateIndustries(request models.GetCandidateIndustriesRequest) (models.CandidateIndustriesSummary, error)
	// Provides total contributed to a candidate from an industry.
	// https://www.opensecrets.org/api/?method=candIndByInd&output=doc
	GetCandidateIndustryDetails(request models.GetCandidateIndustryDetailsRequest) (models.CandidateIndustryDetails, error)
	// Provides sector total of a candidate's receipts
	// https://www.opensecrets.org/api/?method=candSector&output=doc
	GetCandidateTopSectorDetails(request models.GetCandidateTopSectorsRequest) (models.CandidateTopSectorDetails, error)
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

// Construct an OpenSecretsClient with the provided API key and a default http.Client (with a timeout of 5 seconds).
func NewOpenSecretsClient(apikey string) OpenSecretsClient {
	return &openSecretsClient{apiKey: apikey, client: &http.Client{Timeout: time.Second * 5}, validator: validator.New()}
}

// Construct an OpenSecretsClient with the provided API key and a custom HTTP client.
func NewOpenSecretsClientWithHttpClient(apikey string, client OpenSecretsHttpClient) OpenSecretsClient {
	return &openSecretsClient{apiKey: apikey, client: client, validator: validator.New()}
}

func (o *openSecretsClient) GetLegislators(request models.GetLegislatorsRequest) ([]models.Legislator, error) {

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

func (o *openSecretsClient) GetMemberPFDProfile(request models.GetMemberPFDRequest) (models.MemberProfile, error) {
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

func (o *openSecretsClient) GetCandidateSummary(request models.GetCandidateSummaryRequest) (models.CandidateSummary, error) {
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

func (o *openSecretsClient) GetCandidateContributors(request models.GetCandidateContributorsRequest) (models.CandidateContributorSummary, error) {
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

func (o *openSecretsClient) GetCandidateIndustries(request models.GetCandidateIndustriesRequest) (models.CandidateIndustriesSummary, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return models.CandidateIndustriesSummary{}, err
	}

	url := buildGetCandidateIndustriesURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return models.CandidateIndustriesSummary{}, err
	}

	return parse.ParseCandidateIndustriesJSON(responseBody)
}

func (o *openSecretsClient) GetCandidateIndustryDetails(request models.GetCandidateIndustryDetailsRequest) (models.CandidateIndustryDetails, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return models.CandidateIndustryDetails{}, err
	}

	url := buildGetCandidateIndustryDetailsURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return models.CandidateIndustryDetails{}, err
	}

	return parse.ParseCandidateIndustryDetailsJSON(responseBody)
}

func (o *openSecretsClient) GetCandidateTopSectorDetails(request models.GetCandidateTopSectorsRequest) (models.CandidateTopSectorDetails, error) {
	err := o.validator.Struct(request)

	if err != nil {
		return models.CandidateTopSectorDetails{}, err
	}

	url := buildGetCandidatTopSectorsURL(request, o.apiKey)

	responseBody, err := o.makeGETRequest(url)

	if err != nil {
		return models.CandidateTopSectorDetails{}, err
	}

	return parse.ParseCandidateTopSectorsJSON(responseBody)
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
