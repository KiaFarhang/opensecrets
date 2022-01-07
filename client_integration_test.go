package opensecrets

import (
	"net/http"
	"os"
	"testing"
	"time"
)

const no_api_key_error_message string = "You must provide an API_KEY environment variable for end-to-end tests. To just run unit tests, pass the -short flag to the go test command."

func TestClientEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Fatal(no_api_key_error_message)
	}

	httpClient := &http.Client{Timeout: time.Second * 5}
	client := NewOpenSecretsClientWithHttpClient(apiKey, httpClient)

	t.Run("GetLegislators", func(t *testing.T) {
		request := GetLegislatorsRequest{Id: "TX"}
		legislators, err := client.GetLegislators(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetLegislators endpoint", err.Error())
		}

		if len(legislators) == 0 {
			t.Fatalf("Got 0 legislators from GetLegislators call")
		}

	})

	t.Run("GetMemberPFDProfile", func(t *testing.T) {
		request := GetMemberPFDRequest{Cid: "N00007360", Year: 2016}

		memberProfile, err := client.GetMemberPFDProfile(request)

		if err != nil {
			t.Fatalf("Got error %s calling GetMemberPFDProfile", err.Error())
		}

		memberName := memberProfile.Name
		wantedName := "Pelosi, Nancy"

		assertStringMatches(memberName, wantedName, t)

		memberAssets := memberProfile.Assets

		assertIntMatches(len(memberAssets), 5, t)

		memberTransactions := memberProfile.Transactions

		assertIntMatches(len(memberTransactions), 5, t)

		memberPositions := memberProfile.Positions

		assertIntMatches(len(memberPositions), 5, t)

	})

	t.Run("GetCandidateSummary", func(t *testing.T) {
		request := GetCandidateSummaryRequest{Cid: "N00007360", Cycle: 2022}
		candidateSummary, err := client.GetCandidateSummary(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetCandidateSummary endpoint", err.Error())
		}

		assertStringMatches(candidateSummary.CandidateName, "Pelosi, Nancy", t)
		assertIntMatches(candidateSummary.Cycle, 2022, t)
		assertIntMatches(candidateSummary.FirstElected, 1987, t)
	})

}
