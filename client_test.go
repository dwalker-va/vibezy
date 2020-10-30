package vibezy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_SetsRequestHeaders(t *testing.T) {
	// arrange
	c := NewClient("t3stT0K3n")

	// act
	r, err := c.buildRequest(context.Background(), http.MethodGet, "path", nil)
	if err != nil {
		t.Fatalf("received unexpected error: %s", err.Error())
	}

	// assert
	want := "application/json"
	got := r.Header.Get("content-type")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("header mismatch (-want +got):\n%s", diff)
	}

	want = "Bearer t3stT0K3n"
	got = r.Header.Get("Authorization")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("header mismatch (-want +got):\n%s", diff)
	}
}

func TestClient_SuccessfulAPICall(t *testing.T) {
	// arrange
	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				b, _ := json.Marshal(PingResponse{
					IsSuccess:    true,
					ErrorMessage: "",
				})
				_, _ = fmt.Fprint(w, string(b))
			}),
	)
	c := NewClient("t3stT0K3n")
	c.baseURL = s.URL

	// act
	resp, err := c.Ping(context.Background())

	// assert
	if diff := cmp.Diff(nil, err); diff != "" {
		t.Fatalf("error mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(true, resp.IsSuccess); diff != "" {
		t.Fatalf("resp.IsSuccess mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff("", resp.ErrorMessage); diff != "" {
		t.Fatalf("resp.ErrorMessage mismatch (-want +got):\n%s", diff)
	}
}

func TestClient_NonJSONResponse(t *testing.T) {
	// arrange
	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = fmt.Fprint(w, "<html><body>Login</body></html>")
			}),
	)
	c := NewClient("t3stT0K3n")
	c.baseURL = s.URL

	// act
	resp, err := c.Ping(context.Background())

	// assert
	if diff := cmp.Diff("could not decode OfficeVibe response, have you tested whether your API key is set up correctly?: https://api.officevibe.com/docs/ping, invalid character '<' looking for beginning of value", err.Error()); diff != "" {
		t.Fatalf("error mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff((*PingResponse)(nil), resp); diff != "" {
		t.Fatalf("resp mismatch (-want +got):\n%s", diff)
	}
}

func TestClient_UnsuccessfulResponse(t *testing.T) {
	// arrange
	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				b, _ := json.Marshal(PingResponse{
					IsSuccess:    false,
					ErrorMessage: "error from OfficeVibe's server",
				})
				_, _ = fmt.Fprint(w, string(b))
			}),
	)
	c := NewClient("t3stT0K3n")
	c.baseURL = s.URL

	// act
	resp, err := c.Ping(context.Background())

	// assert
	if diff := cmp.Diff("OfficeVibe error: status: `200`, message: `error from OfficeVibe's server`", err.Error()); diff != "" {
		t.Fatalf("error mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(&PingResponse{
		IsSuccess:    false,
		ErrorMessage: "error from OfficeVibe's server",
	}, resp); diff != "" {
		t.Fatalf("resp mismatch (-want +got):\n%s", diff)
	}
}
