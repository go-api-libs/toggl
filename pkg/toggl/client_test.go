package toggl_test

import (
	"testing"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

func TestNewClientWithAPIToken(t *testing.T) {
	if _, err := toggl.NewClientWithAPIToken(""); err == nil {
		t.Fatal("expected error")
	} else if "username is empty" != err.Error() {
		t.Fatalf("want: username is empty, got: %v", err)
	}

	if _, err := toggl.NewClientWithAPIToken("my api token"); err != nil {
		t.Fatal(err)
	}
}
