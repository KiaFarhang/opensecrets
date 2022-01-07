package opensecrets

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseGetLegislatorsJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json := []byte(`{"response": {"legislator": [{"@attributes": {"first_elected": "2000"}}]}}`)
		legislators, err := parseGetLegislatorsJSON(json)
		if err != nil {
			t.Fatalf("Expected no error but got one with message %s", err.Error())
		}
		expectedLegislators := []Legislator{
			{FirstElected: 2000},
		}

		if !reflect.DeepEqual(legislators, expectedLegislators) {
			t.Fatalf("Got %v want %v", legislators, expectedLegislators)
		}
	})
	t.Run("Returns an error for invalid JSON", func(t *testing.T) {
		json := []byte(`GARBAGE`)
		_, err := parseGetLegislatorsJSON(json)
		wantedErrorMessage := unable_to_parse_error_message
		assertErrorMessage(err, wantedErrorMessage, t)
	})
}

func TestParseMemberPFDJSON(t *testing.T) {
	t.Run("Correctly parses valid JSON", func(t *testing.T) {
		json, err := ioutil.ReadFile("mocks/mockPFDResponse.json")
		if err != nil {
			t.Fatalf("Error reading mock data from file: %s", err.Error())
		}
		member, err := parseMemberPFDJSON(json)
		if err != nil {
			t.Fatalf("Expected no error but got one with message %s", err.Error())
		}

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
		wantedErrorMessage := unable_to_parse_error_message
		assertErrorMessage(err, wantedErrorMessage, t)
	})
}
