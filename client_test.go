package vibezy

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestClient_Ping(t *testing.T) {
	cases := []struct {
		apiKey      string
		description string
		want        error
	}{
		{
			apiKey:      os.Getenv("OFFICEVIBE_API_KEY"),
			description: "Returns no error with working API Key",
			want:        nil,
		},
		{
			apiKey:      "brokenApiKey",
			description: "Returns expected error with broken API Key",
			want:        errors.New("could not decode OfficeVibe response, have you tested whether your API key is set up correctly?: https://api.officevibe.com/docs/ping"),
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			client := NewClient(c.apiKey)
			got := client.Ping(context.Background())
			if diff := cmp.Diff(c.want, got, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("Client.Ping() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
