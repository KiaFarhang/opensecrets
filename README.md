# OpenSecrets API Go Client

This is a Go client for the [OpenSecrets campaign finance API.](https://www.opensecrets.org/open-data/api) The following methods are currently supported:

| **API method**    | **Function**             | **Status** |
|-------------------|--------------------------|--------|
| getLegislators    | GetLegislators           | :heavy_check_mark:      |
| memPFDProfile     | GetMemberPFDProfile      | :heavy_check_mark:      |
| candSummary       | GetCandidateSummary      | :heavy_check_mark:      |
| candContrib       | GetCandidateContributors | :heavy_check_mark:      |
| candIndustry      | GetCandidateIndustries   | :heavy_check_mark:      |
| candIndByInd      |                          | TODO   |
| candSector        |                          | TODO   |
| congCmteIndus     |                          | TODO   |
| getOrgs           |                          | TODO   |
| orgSummary        |                          | TODO   |
| independentExpend |                          | TODO   |

[![Go Report Card](https://goreportcard.com/badge/github.com/KiaFarhang/opensecrets)](https://goreportcard.com/report/github.com/KiaFarhang/opensecrets)

## Installing

`go get github.com/KiaFarhang/opensecrets`

## Running tests

Run unit tests with `go test -short ./...`

Run unit and end-to-end tests with `API_KEY=foo go test ./...`