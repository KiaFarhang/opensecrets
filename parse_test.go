package opensecrets

import (
	"io/ioutil"
	"testing"
)

func TestParseGetLegislatorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json := []byte(`{"response": {"legislator": [{"@attributes": {"first_elected": "2000"}}]}}`)
		legislators, err := parseGetLegislatorsJSON(json)
		assertNoError(err, t)

		assertSliceLength(len(legislators), 1, t)

		leigslator := legislators[0]

		assertIntMatches(leigslator.FirstElected, 2000, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := parseGetLegislatorsJSON(json)
		assertErrorMessage(err, unable_to_parse_error_message, t)
	})
}

func TestParseMemberPFDJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("mocks/mockPFDResponse.json")
		assertNoError(err, t)

		member, err := parseMemberPFDJSON(json)
		assertNoError(err, t)

		expectedName := "Pelosi, Nancy"

		if member.Name != expectedName {
			t.Fatalf("Got name %s want %s", member.Name, expectedName)
		}

		assertSliceLength(len(member.Assets), 1, t)

		asset := member.Assets[0]
		wantedAssetName := "25 Point Lobos - Commercial Property"

		assertStringMatches(asset.Name, wantedAssetName, t)

		assertSliceLength(len(member.Transactions), 1, t)

		transaction := member.Transactions[0]
		wantedTransactionAction := "Purchased"

		assertStringMatches(transaction.TransactionAction, wantedTransactionAction, t)

		assertSliceLength(len(member.Positions), 1, t)

		position := member.Positions[0]
		wantedPositionTitle := "Honorary Advisory Board"

		assertStringMatches(position.Title, wantedPositionTitle, t)

	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := parseMemberPFDJSON(json)
		assertErrorMessage(err, unable_to_parse_error_message, t)
	})
}

func TestParseCandidateSummaryJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("mocks/mockCandidateSummaryResponse.json")
		assertNoError(err, t)

		candidateSummary, err := parseCandidateSummaryJSON(json)
		assertNoError(err, t)

		expectedName := "Pelosi, Nancy"
		assertStringMatches(candidateSummary.CandidateName, expectedName, t)

		expectedTotal := 9235427.16
		if candidateSummary.Total != expectedTotal {
			t.Errorf("Wanted %f got %f", expectedTotal, candidateSummary.Total)
		}

	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := parseCandidateSummaryJSON(json)
		assertErrorMessage(err, unable_to_parse_error_message, t)
	})
}

func TestParseCandidateContributorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("mocks/mockCandidateContributorsResponse.json")
		assertNoError(err, t)

		contributorSummary, err := parseCandidateContributorsJSON(json)
		assertNoError(err, t)

		expectedName := "Nancy Pelosi (D)"

		assertStringMatches(contributorSummary.CandidateName, expectedName, t)

		contributors := contributorSummary.Contributors

		assertSliceLength(len(contributors), 10, t)

		firstContributor := contributors[0]

		expectedFirstContributorName := "University of California"

		assertStringMatches(firstContributor.OrganizationName, expectedFirstContributorName, t)

		expectedFirstContributorTotal := float64(130682)

		if firstContributor.Total != expectedFirstContributorTotal {
			t.Errorf("Got float %f wanted %f", firstContributor.Total, expectedFirstContributorTotal)
		}

	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := parseCandidateContributorsJSON(json)
		assertErrorMessage(err, unable_to_parse_error_message, t)
	})
}
