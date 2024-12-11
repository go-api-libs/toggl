package main

import "net/http"

const serverURL = "https://api.track.toggl.com/api/v9"

// probe calls the API server to check what we can do
func probe() error {
	// In general, for HTTP Basic Auth, you have to add the Authorization header with the request. The Authorization header is constructed as follows:

	// In case email and password are used, they are combined into a email:password format
	// In case the api token is used, it is combined in xxxx:api_token format (xxxx indicating user's personal token)
	// The resulting string literal is then encoded using Base64
	// The authorization method and a space i.e. "Basic " is then put before the encoded string.

	req, err := http.NewRequest(http.MethodGet, serverURL+"/me", nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	return err
}
