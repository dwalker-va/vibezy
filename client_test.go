package vibezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_SetsRequestHeaders(t *testing.T) {
	// arrange
	c := NewClient("t3stT0K3n")

	// act
	r, err := c.buildRequest(context.Background(), http.MethodGet, "path", nil)
	if err != nil {
		t.Errorf("received unexpected error: %s", err.Error())
	}

	// assert
	want := "application/json"
	got := r.Header.Get("content-type")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("header mismatch (-want +got):\n%s", diff)
	}

	want = "Bearer t3stT0K3n"
	got = r.Header.Get("Authorization")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("header mismatch (-want +got):\n%s", diff)
	}
}
