package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const BASE_URL string = "http://www.opensecrets.org/api/"

func main() {
	client := &http.Client{}
	apiKey := os.Getenv("API_KEY")
	url := BASE_URL + "?method=getLegislators&output=json&id=TX&apikey=" + apiKey

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error constructing HTTP request: %s", err)
	}

	// The API blocks requests without a user agent
	request.Header.Set("User-Agent", "Golang")

	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("Error calling API: %s", err)
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading contents of response body: %s", err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}
