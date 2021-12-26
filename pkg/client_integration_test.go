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
	client := OpenSecretsClient{httpClient, apiKey}

	t.Run("GetLegislators", func(t *testing.T) {
		request := GetLegislatorsRequest{id: "TX"}
		legislators, err := client.GetLegislators(request)
		if err != nil {
			t.Fatalf("Got error %s calling GetLegislators endpoint", err.Error())
		}
		t.Logf("count of legislators: %d", len(legislators))
		for _, legislator := range legislators {
			t.Logf("legislator first elected: %d", legislator.FirstElected)
		}
	})

}
