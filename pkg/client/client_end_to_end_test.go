package client

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/KiaFarhang/opensecrets/internal/test"
	"github.com/KiaFarhang/opensecrets/pkg/models"
)

const noApiKeyErrorMessage string = "You must provide an API_KEY environment variable for end-to-end tests. To just run unit tests, pass the -short flag to the go test command."

func TestClientEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Fatal(noApiKeyErrorMessage)
	}

	httpClient := &http.Client{Timeout: time.Second * 5}
	client := NewOpenSecretsClientWithHttpClient(apiKey, httpClient)

	t.Run("GetLegislators", func(t *testing.T) {
		request := models.LegislatorsRequest{Id: "TX"}
		legislators, err := client.GetLegislators(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetLegislators", err.Error())
		}

		if len(legislators) == 0 {
			t.Fatalf("Got 0 legislators from GetLegislators call")
		}

	})

	t.Run("GetMemberPFDProfile", func(t *testing.T) {
		request := models.MemberPFDRequest{Cid: "N00007360", Year: 2016}

		memberProfile, err := client.GetMemberPFDProfile(request)

		if err != nil {
			t.Fatalf("Got error %s calling GetMemberPFDProfile", err.Error())
		}

		memberName := memberProfile.Name
		wantedName := "Pelosi, Nancy"

		test.AssertStringMatches(memberName, wantedName, t)

		memberAssets := memberProfile.Assets

		test.AssertIntMatches(len(memberAssets), 5, t)

		memberTransactions := memberProfile.Transactions

		test.AssertIntMatches(len(memberTransactions), 5, t)

		memberPositions := memberProfile.Positions

		test.AssertIntMatches(len(memberPositions), 5, t)

	})

	t.Run("GetCandidateSummary", func(t *testing.T) {
		request := models.CandidateSummaryRequest{Cid: "N00007360", Cycle: 2022}
		candidateSummary, err := client.GetCandidateSummary(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetCandidateSummary", err.Error())
		}

		test.AssertStringMatches(candidateSummary.CandidateName, "Pelosi, Nancy", t)
		test.AssertIntMatches(candidateSummary.Cycle, 2022, t)
		test.AssertIntMatches(candidateSummary.FirstElected, 1987, t)
	})

	t.Run("GetCandidateContributors", func(t *testing.T) {
		request := models.CandidateContributorsRequest{Cid: "N00007360", Cycle: 2018}
		candidateContributorSummary, err := client.GetCandidateContributors(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetCandidateContributors", err.Error())
		}

		test.AssertStringMatches(candidateContributorSummary.CandidateName, "Nancy Pelosi (D)", t)
		test.AssertSliceLength(len(candidateContributorSummary.Contributors), 10, t)
	})

	t.Run("GetCandidateIndustries", func(t *testing.T) {
		request := models.CandidateIndustriesRequest{Cid: "N00005681", Cycle: 2018}
		summary, err := client.GetCandidateIndustries(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetCandidateIndustries", err.Error())
		}

		test.AssertStringMatches(summary.CandidateName, "Pete Sessions (R)", t)
		test.AssertSliceLength(len(summary.Industries), 10, t)
	})

	t.Run("GetCandidateIndustryDetails", func(t *testing.T) {
		request := models.CandidateIndustryDetailsRequest{Cid: "N00007360", Ind: "K02", Cycle: 2020}
		details, err := client.GetCandidateIndustryDetails(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetCandidateIndustryDetails", err.Error())
		}

		test.AssertStringMatches(details.Chamber, "H", t)
		test.AssertNotZero(details.Total, t)
	})

	t.Run("GetCandidateTopSectorDetails", func(t *testing.T) {
		request := models.CandidateTopSectorsRequest{Cid: "N00007360", Cycle: 2020}
		details, err := client.GetCandidateTopSectorDetails(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetCandidateTopSectorDetails", err.Error())
		}

		test.AssertStringMatches(details.CandidateName, "Nancy Pelosi (D)", t)
		test.AssertSliceLength(len(details.Sectors), 13, t)
		firstSector := details.Sectors[0]

		expectedSectorId := "A"
		test.AssertStringMatches(firstSector.Id, expectedSectorId, t)

		test.AssertNotZero(firstSector.Individuals, t)
	})

	t.Run("GetCommitteeFundraisingDetails", func(t *testing.T) {
		request := models.FundraisingByCongressionalCommitteeRequest{Committee: "HARM", Industry: "F10", CongressNumber: 116}
		details, err := client.GetCommitteeFundraisingDetails(request)
		if err != nil {
			t.Fatalf("Got error %s when calling GetCommitteeFundraisingDetails", err.Error())
		}

		test.AssertStringMatches(details.CommitteeName, "HARM", t)
		test.AssertSliceLength(len(details.Members), 56, t)

		firstMember := details.Members[0]
		test.AssertStringMatches(firstMember.State, "New York", t)
		test.AssertNotZero(firstMember.Pacs, t)
	})

	t.Run("SearchForOrganization", func(t *testing.T) {
		request := models.OrganizationSearch{Name: "Goldman"}
		searchResults, err := client.SearchForOrganization(request)

		if err != nil {
			t.Fatalf("Got error %s when calling SearchForOrganization", err.Error())
		}

		if len(searchResults) == 0 {
			t.Fatal("Got 0 results from SearchForOrganization call")
		}
	})

	t.Run("GetOrganizationSummary", func(t *testing.T) {
		request := models.OrganizationSummaryRequest{Id: "D000000125"}
		summary, err := client.GetOrganizationSummary(request)

		if err != nil {
			t.Fatalf("Got error %s when calling GetOrganizationSummary", err.Error())
		}

		test.AssertStringMatches(summary.Name, "General Electric", t)
	})
}
