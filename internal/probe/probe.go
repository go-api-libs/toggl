package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
)

const serverURL = "https://api.track.toggl.com/api/v9"

// probe calls the API server to check what we can do
func probe() error {
	req, err := http.NewRequest(http.MethodGet, serverURL+"/me", nil)
	if err != nil {
		return err
	}

	tkn := os.Getenv("TOGGL_TOKEN")

	req.SetBasicAuth(tkn, "api_token")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("got %s", resp.Status)
	}

	return nil
}

func maskSecrets(i *cassette.Interaction) error {
	tkn := os.Getenv("TOGGL_TOKEN")
	i.Response.Body = strings.ReplaceAll(i.Response.Body,
		tkn, strings.Repeat("*", len(tkn)))

	return nil
}
