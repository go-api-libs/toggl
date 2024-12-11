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
	fmt.Printf("tkn: %v\n", tkn)

	req.SetBasicAuth(tkn, "api_token")

	// In general, for HTTP Basic Auth, you have to add the Authorization header with the request. The Authorization header is constructed as follows:

	// In case email and password are used, they are combined into a email:password format
	// In case the api token is used, it is combined in xxxx:api_token format (xxxx indicating user's personal token)
	// The resulting string literal is then encoded using Base64

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
