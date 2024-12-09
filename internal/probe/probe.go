package main

import "net/http"

const serverURL = "https://api.track.toggl.com/api/v9"

// probe calls the API server to check what we can do
func probe() error {
	_, err := http.Get(serverURL + "/me")
	return err
}
