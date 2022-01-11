package parse

import (
	"io/ioutil"
	"testing"

	"github.com/KiaFarhang/opensecrets/internal/test"
)

func TestParseGetLegislatorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json := []byte(`{"response": {"legislator": [{"@attributes": {"first_elected": "2000"}}]}}`)
		legislators, err := ParseGetLegislatorsJSON(json)
		test.AssertNoError(err, t)

		test.AssertSliceLength(len(legislators), 1, t)

		leigslator := legislators[0]

		test.AssertIntMatches(leigslator.FirstElected, 2000, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseGetLegislatorsJSON(json)
		test.AssertErrorMessage(err, Unable_to_parse_error_message, t)
	})
}

func TestParseMemberPFDJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockPFDResponse.json")
		test.AssertNoError(err, t)

		member, err := ParseMemberPFDJSON(json)
		test.AssertNoError(err, t)

		expectedName := "Pelosi, Nancy"

		if member.Name != expectedName {
			t.Fatalf("Got name %s want %s", member.Name, expectedName)
		}

		test.AssertSliceLength(len(member.Assets), 1, t)

		asset := member.Assets[0]
		wantedAssetName := "25 Point Lobos - Commercial Property"

		test.AssertStringMatches(asset.Name, wantedAssetName, t)

		test.AssertSliceLength(len(member.Transactions), 1, t)

		transaction := member.Transactions[0]
		wantedTransactionAction := "Purchased"

		test.AssertStringMatches(transaction.TransactionAction, wantedTransactionAction, t)

		test.AssertSliceLength(len(member.Positions), 1, t)

		position := member.Positions[0]
		wantedPositionTitle := "Honorary Advisory Board"

		test.AssertStringMatches(position.Title, wantedPositionTitle, t)

	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseMemberPFDJSON(json)
		test.AssertErrorMessage(err, Unable_to_parse_error_message, t)
	})
}

func TestParseCandidateSummaryJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockCandidateSummaryResponse.json")
		test.AssertNoError(err, t)

		candidateSummary, err := ParseCandidateSummaryJSON(json)
		test.AssertNoError(err, t)

		expectedName := "Pelosi, Nancy"
		test.AssertStringMatches(candidateSummary.CandidateName, expectedName, t)

		expectedTotal := 9235427.16
		if candidateSummary.Total != expectedTotal {
			t.Errorf("Wanted %f got %f", expectedTotal, candidateSummary.Total)
		}

	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseCandidateSummaryJSON(json)
		test.AssertErrorMessage(err, Unable_to_parse_error_message, t)
	})
}

func TestParseCandidateContributorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockCandidateContributorsResponse.json")
		test.AssertNoError(err, t)

		contributorSummary, err := ParseCandidateContributorsJSON(json)
		test.AssertNoError(err, t)

		expectedName := "Nancy Pelosi (D)"

		test.AssertStringMatches(contributorSummary.CandidateName, expectedName, t)

		contributors := contributorSummary.Contributors

		test.AssertSliceLength(len(contributors), 10, t)

		firstContributor := contributors[0]

		expectedFirstContributorName := "University of California"

		test.AssertStringMatches(firstContributor.OrganizationName, expectedFirstContributorName, t)

		expectedFirstContributorTotal := float64(130682)

		if firstContributor.Total != expectedFirstContributorTotal {
			t.Errorf("Got float %f wanted %f", firstContributor.Total, expectedFirstContributorTotal)
		}

	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseCandidateContributorsJSON(json)
		test.AssertErrorMessage(err, Unable_to_parse_error_message, t)
	})
}