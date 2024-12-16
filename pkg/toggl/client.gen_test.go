// This file provides tests for the client package.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package toggl_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/go-api-libs/api"
	"github.com/go-api-libs/toggl/pkg/toggl"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

type testRoundTripper struct {
	rsp *http.Response
	err error
}

func (t *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.rsp, t.err
}

func TestClient_Error(t *testing.T) {
	ctx := context.Background()

	if _, err := toggl.NewClient("", ""); err == nil {
		t.Fatal("expected error")
	} else if "username is empty" != err.Error() {
		t.Fatalf("want: username is empty, got: %v", err)
	}

	if _, err := toggl.NewClient("myUsername", ""); err == nil {
		t.Fatal("expected error")
	} else if "password is empty" != err.Error() {
		t.Fatalf("want: password is empty, got: %v", err)
	}

	c, err := toggl.NewClient("myUsername", "myPassword")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Do", func(t *testing.T) {
		testErr := errors.New("test error")
		http.DefaultClient.Transport = &testRoundTripper{err: testErr}

		if _, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true}); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, testErr) {
			t.Fatalf("want: %v, got: %v", testErr, err)
		}

		if _, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
			CreatedWith: "github.com/go-api-libs/toggl",
			Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
			WorkspaceID: 2230580,
		}); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, testErr) {
			t.Fatalf("want: %v, got: %v", testErr, err)
		}

		if _, err := c.GetCurrentTimeEntry(ctx); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, testErr) {
			t.Fatalf("want: %v, got: %v", testErr, err)
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		errDecode := &api.DecodingError{}

		t.Run("GetMe", func(t *testing.T) {
			// unknown status code
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{StatusCode: http.StatusTeapot}}

			if _, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true}); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownStatusCode) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownStatusCode, err)
			}

			// unknown content type for 200 OK
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Header:     http.Header{"Content-Type": []string{"foo"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true}); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownContentType) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownContentType, err)
			}

			// decoding error for known content type "application/json"
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Body:       io.NopCloser(strings.NewReader("{")),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true}); err == nil {
				t.Fatal("expected error")
			} else if !errors.As(err, &errDecode) {
				t.Fatalf("want: %v, got: %v", errDecode, err)
			}
		})

		t.Run("CreateTimeEntry", func(t *testing.T) {
			// unknown status code
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{StatusCode: http.StatusTeapot}}

			if _, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
				WorkspaceID: 2230580,
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownStatusCode) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownStatusCode, err)
			}

			// unknown content type for 200 OK
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Header:     http.Header{"Content-Type": []string{"foo"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
				WorkspaceID: 2230580,
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownContentType) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownContentType, err)
			}

			// decoding error for known content type "application/json"
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Body:       io.NopCloser(strings.NewReader("{")),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
				WorkspaceID: 2230580,
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.As(err, &errDecode) {
				t.Fatalf("want: %v, got: %v", errDecode, err)
			}

			// unknown content type for 400 Bad Request
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Header:     http.Header{"Content-Type": []string{"foo"}},
				StatusCode: http.StatusBadRequest,
			}}

			if _, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
				WorkspaceID: 2230580,
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownContentType) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownContentType, err)
			}

			// decoding error for known content type "application/json"
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Body:       io.NopCloser(strings.NewReader("{")),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				StatusCode: http.StatusBadRequest,
			}}

			if _, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
				WorkspaceID: 2230580,
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.As(err, &errDecode) {
				t.Fatalf("want: %v, got: %v", errDecode, err)
			}
		})

		t.Run("GetCurrentTimeEntry", func(t *testing.T) {
			// unknown status code
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{StatusCode: http.StatusTeapot}}

			if _, err := c.GetCurrentTimeEntry(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownStatusCode) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownStatusCode, err)
			}

			// unknown content type for 200 OK
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Header:     http.Header{"Content-Type": []string{"foo"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetCurrentTimeEntry(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownContentType) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownContentType, err)
			}

			// decoding error for known content type "application/json"
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Body:       io.NopCloser(strings.NewReader("{")),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetCurrentTimeEntry(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.As(err, &errDecode) {
				t.Fatalf("want: %v, got: %v", errDecode, err)
			}
		})
	})
}

func replay(t *testing.T, cassette string) {
	t.Helper()

	r, err := recorder.NewWithOptions(&recorder.Options{
		CassetteName:       cassette,
		Mode:               recorder.ModeReplayOnly,
		RealTransport:      http.DefaultTransport,
		SkipRequestLatency: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = r.Stop()
	})

	r.SetMatcher(matcher)
	http.DefaultClient.Transport = r
}

func matcher(r *http.Request, i cassette.Request) bool {
	if !cassette.DefaultMatcher(r, i) {
		return false
	}

	return getBody(r) == i.Body
}

func getBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}

	if r.GetBody == nil {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewReader(b))
		return string(b)
	}

	body, err := r.GetBody()
	if err != nil {
		panic(err)
	}
	defer body.Close()

	b, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func TestClient_VCR(t *testing.T) {
	ctx := context.Background()

	c, err := toggl.NewClient("myUsername", "myPassword")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("2024-12-09", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-09")

		if _, err := c.GetMe(ctx, nil); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, api.ErrStatusCode) {
			t.Fatalf("want: %v, got: %v", api.ErrStatusCode, err)
		}
	})

	t.Run("2024-12-10", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-10")

		if _, err := c.GetMe(ctx, nil); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, api.ErrStatusCode) {
			t.Fatalf("want: %v, got: %v", api.ErrStatusCode, err)
		}
	})

	t.Run("2024-12-11", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-11")

		res, err := c.GetMe(ctx, nil)
		if err != nil {
			t.Fatal(err)
		} else if res == nil {
			t.Fatal("result is nil")
		}
	})

	t.Run("2024-12-12", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-12")

		{
			res, err := c.GetMe(ctx, nil)
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}
	})

	t.Run("2024-12-13", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-13")

		{
			res, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.GetCurrentTimeEntry(ctx)
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}
	})

	t.Run("2024-12-14", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-14")

		{
			res, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.GetCurrentTimeEntry(ctx)
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}
	})

	t.Run("2024-12-15", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-15")

		{
			res, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.GetCurrentTimeEntry(ctx)
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
				WorkspaceID: 2230580,
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Start:       time.Date(2024, time.December, 15, 21, 19, 39, 215084000, time.Local),
				WorkspaceID: 2230580,
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}
	})

	t.Run("2024-12-16", func(t *testing.T) {
		replay(t, "../../pkg/toggl/vcr/2024-12-16")

		{
			res, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.GetCurrentTimeEntry(ctx)
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Duration:    time.Hour,
				Start:       time.Date(2024, time.December, 16, 3, 18, 20, 257412000, time.Local),
				WorkspaceID: 2230580,
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Description: "Hello Toggl",
				Start:       time.Date(2024, time.December, 16, 2, 29, 40, 335227000, time.Local),
				WorkspaceID: 2230580,
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
				CreatedWith: "github.com/go-api-libs/toggl",
				Description: "Hello Toggl",
				Duration:    -time.Second,
				Start:       time.Date(2024, time.December, 16, 2, 31, 14, 86355000, time.Local),
				WorkspaceID: 2230580,
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}
	})
}
