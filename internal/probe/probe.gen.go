package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func main() {
	start := time.Now()
	defer func() { log.Printf("Probe took %s", time.Since(start)) }()

	r, err := record()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// Ensure recorder is saved once we're done
		if err := r.Stop(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := probe(); err != nil {
		log.Fatal(err)
	}
}

func record() (*recorder.Recorder, error) {
	r, err := recorder.NewWithOptions(&recorder.Options{
		CassetteName:       fmt.Sprintf("pkg/toggl/vcr/%s", time.Now().Format(time.DateOnly)),
		Mode:               recorder.ModeReplayWithNewEpisodes,
		SkipRequestLatency: true,
		RealTransport:      http.DefaultTransport,
	})
	if err != nil {
		return nil, err
	}

	r.AddHook(maskAuthorization, recorder.BeforeSaveHook)
	r.AddHook(maskSecrets, recorder.BeforeSaveHook)

	// replace the default transport with the recorder
	http.DefaultClient.Transport = r

	return r, nil
}

func maskAuthorization(i *cassette.Interaction) error {
	auth := i.Request.Headers.Get("Authorization")
	if auth == "" {
		return nil
	}

	i.Request.Headers.Set("Authorization", maskAuthString(auth))

	return nil
}

func maskAuthString(s string) string {
	if parts := strings.Split(s, " "); len(parts) == 2 {
		// mask the token except the type (Bearer, Basic, etc.)
		return fmt.Sprintf("%s %s", parts[0], strings.Repeat("*", len(parts[1])))
	}

	return strings.Repeat("*", len(s))
}
