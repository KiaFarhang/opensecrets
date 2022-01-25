# OpenSecrets API Go Client

[![Go Report Card](https://goreportcard.com/badge/github.com/KiaFarhang/opensecrets)](https://goreportcard.com/report/github.com/KiaFarhang/opensecrets)

This is a Go client for the [OpenSecrets campaign finance API.](https://www.opensecrets.org/open-data/api) The following methods are currently supported:

| **API method**    | **Function**             | **Status** |
|-------------------|--------------------------|--------|
| getLegislators    | GetLegislators           | :heavy_check_mark:      |
| memPFDProfile     | GetMemberPFDProfile      | :heavy_check_mark:      |
| candSummary       | GetCandidateSummary      | :heavy_check_mark:      |
| candContrib       | GetCandidateContributors | :heavy_check_mark:      |
| candIndustry      | GetCandidateIndustries   | :heavy_check_mark:      |
| candIndByInd      | GetCandidateIndustryDetails | :heavy_check_mark:   |
| candSector        | GetCandidateTopSectorDetails | :heavy_check_mark:   |
| congCmteIndus     | GetCommitteeFundraisingDetails | :heavy_check_mark:   |
| getOrgs           | SearchForOrganization    | :heavy_check_mark:   |
| orgSummary        | GetOrganizationSummary   | :heavy_check_mark:   |
| independentExpend |                          | TODO   |

## Installing

`go get github.com/KiaFarhang/opensecrets`

## Running tests

Run unit tests with `go test -short ./...`

Run unit and end-to-end tests with `API_KEY=foo go test ./...`