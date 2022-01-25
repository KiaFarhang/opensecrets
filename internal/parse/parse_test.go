package parse

import (
	"io/ioutil"
	"testing"

	"github.com/KiaFarhang/opensecrets/internal/test"
)

func TestParseLegislatorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json := []byte(`{"response": {"legislator": [{"@attributes": {"first_elected": "2000"}}]}}`)
		legislators, err := ParseLegislatorsJSON(json)
		test.AssertNoError(err, t)

		test.AssertSliceLength(len(legislators), 1, t)

		leigslator := legislators[0]

		test.AssertIntMatches(leigslator.FirstElected, 2000, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseLegislatorsJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
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
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
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
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
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

		test.AssertFloat64Matches(firstContributor.Total, expectedFirstContributorTotal, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseCandidateContributorsJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
	})
}

func TestParseCandidateIndustriesJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockCandidateIndustriesResponse.json")
		test.AssertNoError(err, t)

		industrySummary, err := ParseCandidateIndustriesJSON(json)
		test.AssertNoError(err, t)

		expectedName := "Pete Sessions (R)"
		test.AssertStringMatches(industrySummary.CandidateName, expectedName, t)

		test.AssertSliceLength(len(industrySummary.Industries), 10, t)

		topIndustry := industrySummary.Industries[0]

		expectedIndustryName := "Leadership PACs"
		test.AssertStringMatches(topIndustry.IndustryName, expectedIndustryName, t)

		expectedTotal := float64(312081)
		test.AssertFloat64Matches(topIndustry.Total, expectedTotal, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseCandidateIndustriesJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
	})
}

func TestParseCandidateIndustryDetailsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockCandidateIndustryDetailsResponse.json")
		test.AssertNoError(err, t)

		details, err := ParseCandidateIndustryDetailsJSON(json)
		test.AssertNoError(err, t)

		expectedChamber := "H"
		test.AssertStringMatches(details.Chamber, expectedChamber, t)

		expectedTotal := float64(151248)
		test.AssertFloat64Matches(details.Total, expectedTotal, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseCandidateIndustryDetailsJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
	})
}

func TestParseCandidateTopSectorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockCandidateTopSectorsResponse.json")
		test.AssertNoError(err, t)

		details, err := ParseCandidateTopSectorsJSON(json)
		test.AssertNoError(err, t)

		expectedCandidateName := "Nancy Pelosi (D)"
		test.AssertStringMatches(details.CandidateName, expectedCandidateName, t)

		test.AssertSliceLength(len(details.Sectors), 13, t)

		firstSector := details.Sectors[0]

		expectedSectorId := "A"
		test.AssertStringMatches(firstSector.Id, expectedSectorId, t)

		expectedSectorIndividuals := float64(125816)
		test.AssertFloat64Matches(firstSector.Individuals, expectedSectorIndividuals, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseCandidateTopSectorsJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
	})
}

func TestParseFundraisingByCommitteeJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockFundraisingByCommitteeResponse.json")
		test.AssertNoError(err, t)

		details, err := ParseFundraisingByCommitteeJSON(json)
		test.AssertNoError(err, t)

		test.AssertStringMatches(details.Industry, "Real Estate", t)
		test.AssertIntMatches(details.CongressNumber, 116, t)

		test.AssertSliceLength(len(details.Members), 56, t)

		firstMember := details.Members[0]

		test.AssertStringMatches(firstMember.Name, "Stefanik, Elise", t)
		expectedTotal := float64(402408)
		test.AssertFloat64Matches(firstMember.Total, expectedTotal, t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseFundraisingByCommitteeJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
	})
}

func TestParseOrganizationSearchJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockOrganizationSearchResponse.json")
		test.AssertNoError(err, t)

		searchResults, err := ParseOrganizationSearchJSON(json)
		test.AssertNoError(err, t)

		test.AssertSliceLength(len(searchResults), 10, t)

		firstResult := searchResults[0]

		test.AssertStringMatches(firstResult.Id, "D000070392", t)
		test.AssertStringMatches(firstResult.Name, "Goldman Environmental Prize", t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseOrganizationSearchJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
	})
}

func TestParseOrganizationSummaryJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("../mocks/mockOrganizationSummaryResponse.json")
		test.AssertNoError(err, t)

		summary, err := ParseOrganizationSummaryJSON(json)
		test.AssertNoError(err, t)

		test.AssertStringMatches(summary.Name, "General Electric", t)
		test.AssertFloat64Matches(summary.Soft, float64(2236), t)
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := ParseOrganizationSummaryJSON(json)
		test.AssertErrorMessage(err, UnableToParseErrorMessage, t)
	})
}
