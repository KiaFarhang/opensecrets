# OpenSecrets API Go Client

[![Go Report Card](https://goreportcard.com/badge/github.com/KiaFarhang/opensecrets)](https://goreportcard.com/report/github.com/KiaFarhang/opensecrets)

This is a Go client for the [OpenSecrets campaign finance API.](https://www.opensecrets.org/open-data/api)

## Installation

`go get github.com/KiaFarhang/opensecrets`

## Usage

### Constructing a client

First, instantiate an `OpenSecretsClient` by passing your API key to the constructor function:

```go
package whatever

import (
	"github.com/KiaFarhang/opensecrets/pkg/client"
	"github.com/KiaFarhang/opensecrets/pkg/models"
)

openSecretsClient := client.NewOpenSecretsClient("YOUR_API_KEY")
```

If you'd like to customize the HTTP client the library uses, call the `NewOpenSecretsClientWithHttpClient` constructor instead:

```go
httpClient := &http.Client{Timeout: time.Second * 3} // Whatever other configuration you want here...

openSecretsClient := client.NewOpenSecretsClientWithHttpClient("YOUR_API_KEY", httpClient)
```

The default client has a 5-second HTTP timeout.

The `OpenSecretsClient` is thread safe; you should construct one and share it throughout your application.

### Making API calls

The client has public methods for [each of the OpenSecrets API methods](https://www.opensecrets.org/open-data/api-documentation). To call one, just pass it the appropriate request object from the `models` package:

```go
request := models.LegislatorsRequest{Id: "TX"}
legislators, err := client.GetLegislators(request)
```

The client will either return a struct containing the data from the API call or an error if something went wrong.

The client also throws an error if you pass it a request that's missing a required parameter. See the table below for details on which fields are required.

### Available methods

| API method | Client method | Description | Docs |
|---|---|---|---|
| getLegislators | GetLegislators | Provides a list of Congressional legislators for a specified subset (state or specific CID) | [Link](https://www.opensecrets.org/api/?method=getLegislators&output=doc) |
| memPFDProfile | GetMemberPFDProfile | Returns data on the personal finances of a member of Congress, as well as judicial + executive branches | [Link](https://www.opensecrets.org/api/?method=memPFDprofile&output=doc) |
| candSummary | GetCandidateSummary | Provides summary fundraising information for a politician | [Link](https://www.opensecrets.org/api/?method=candSummary&output=doc) |
| candContrib | GetCandidateContributors | Returns top contributors to a candidate for/sitting member of Congress | [Link](https://www.opensecrets.org/api/?method=candContrib&output=doc) |
| candIndustry | GetCandidateIndustries | Provides the top 10 industries contributing to a candidate | [Link](https://www.opensecrets.org/api/?method=candIndustry&output=doc) |
| candIndByInd | GetCandidateIndustryDetails | Provides total contributed to a candidate from an industry. | [Link](https://www.opensecrets.org/api/?method=candIndByInd&output=doc) |
| candSector | GetCandidateTopSectorDetails | Provides sector total of a candidate's receipts | [Link](https://www.opensecrets.org/api/?method=candSector&output=doc) |
| congCmteIndus | GetCommitteeFundraisingDetails | Provides fundraising details for all members of a given committee from the provided industry | [Link](https://www.opensecrets.org/api/?method=congCmteIndus&output=doc) |
| getOrgs | SearchForOrganization | Searches for an organization by name or partial name | [Link](https://www.opensecrets.org/api/?method=getOrgs&output=doc) |
| orgSummary | GetOrganizationSummary | Provides summary fundraising information for an organization | [Link](https://www.opensecrets.org/api/?method=orgSummary&output=doc) |
| independentExpend |  |  | [Link](https://www.opensecrets.org/api/?method=independentExpend&output=doc) |

## Running tests

Run unit tests with `go test -short ./...`

Run unit and end-to-end tests with `API_KEY=foo go test ./...`