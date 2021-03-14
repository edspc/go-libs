package zoho_test

import (
	"testing"

	"github.com/edspc/go-libs/oauth2/zoho"
	"golang.org/x/oauth2"
)

func TestGetEndpoint(t *testing.T) {
	endpointTests := []struct {
		in  zoho.Datacenter
		out oauth2.Endpoint
	}{
		{
			in: zoho.EUDatacenter,
			out: oauth2.Endpoint{
				AuthURL:   "https://accounts.zoho.eu/oauth/v2/auth",
				TokenURL:  "https://accounts.zoho.eu/oauth/v2/token",
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
		{
			in: zoho.GlobalDatacenter,
			out: oauth2.Endpoint{
				AuthURL:   "https://accounts.zoho.com/oauth/v2/auth",
				TokenURL:  "https://accounts.zoho.com/oauth/v2/token",
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
	}

	for _, tt := range endpointTests {
		t.Run(string(tt.in), func(t *testing.T) {
			endpoint := zoho.GetEndpoint(tt.in)
			if endpoint != tt.out {
				t.Errorf("got %q, want %q", endpoint, tt.out)
			}
		})
	}
}

func TestGetAuthUrl(t *testing.T) {
	conf := zoho.Config{
		Config: &oauth2.Config{
			ClientID:     "c_id",
			ClientSecret: "c_secret",
			RedirectURL:  "http://localhost",
			Scopes:       []string{"Testing"},
			Endpoint:     zoho.GetEndpoint(zoho.EUDatacenter),
		},
	}

	testURL := conf.GetAuthUrl("test")
	expectedURL := "https://accounts.zoho.eu/oauth/v2/auth?access_type=online&client_id=c_id&redirect_uri=http%3A%2F%2Flocalhost&response_type=code&scope=Testing&state=test"

	if testURL != expectedURL {
		t.Errorf("got %q, want %q", testURL, expectedURL)
	}
}
