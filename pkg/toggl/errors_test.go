package toggl

import (
	"encoding/json/v2"
	"errors"
	"testing"
)

func TestErrTimeEntryAlreadyStopped(t *testing.T) {
	var apiErr APIErrorString
	if err := json.Unmarshal([]byte(`"Time entry already stopped"`), &apiErr, jsonOpts); err != nil {
		t.Fatal(err)
	}

	var err error = &apiErr
	if _, ok := errors.AsType[*APIErrorString](err); !ok {
		t.Fatalf("got: %T, want: *APIErrorString", err)
	} else if !errors.Is(err, ErrTimeEntryAlreadyStopped) {
		t.Fatalf("got: %q, want: ErrTimeEntryAlreadyStopped", err)
	}
}

func TestAPIErrorStringIs(t *testing.T) {
	nilp := (*APIErrorString)(nil)
	tests := []struct {
		name        string
		err, target error
		want        bool
	}{
		{"nil,nil", nilp, nilp, true},
		{"nil,sentinel", nilp, ErrTimeEntryAlreadyStopped, false},
		{"sentinel,nil", ErrTimeEntryAlreadyStopped, nilp, false},
		{"sentinel,same", ErrTimeEntryAlreadyStopped, ErrTimeEntryAlreadyStopped, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(tt.err, tt.target); got != tt.want {
				t.Fatalf("errors.Is = %v, want %v", got, tt.want)
			}
		})
	}
}
