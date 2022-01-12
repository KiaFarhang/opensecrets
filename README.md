# OpenSecrets API Go Client

This is a Go client for the [OpenSecrets campaign finance API.](https://www.opensecrets.org/open-data/api) The following methods are currently supported:

| **API method**    | **Function**             | **Status** |
|-------------------|--------------------------|--------|
| getLegislators    | GetLegislators           | ✔️      |
| memPFDProfile     | GetMemberPFDProfile      | ✔️      |
| candSummary       | GetCandidateSummary      | ✔️      |
| candContrib       | GetCandidateContributors | ✔️      |
| candIndustry      |                          | TODO   |
| candIndByInd      |                          | TODO   |
| candSector        |                          | TODO   |
| congCmteIndus     |                          | TODO   |
| getOrgs           |                          | TODO   |
| orgSummary        |                          | TODO   |
| independentExpend |                          | TODO   |

## Installing

`go get github.com/KiaFarhang/opensecrets`

## Running tests

Run unit tests with `go test -short ./...`

Run unit and end-to-end tests with `API_KEY=foo go test ./...`