package zoho

import (
	"context"

	"golang.org/x/oauth2"
)

type Datacenter string

const (
	GlobalDatacenter Datacenter = "com"
	EUDatacenter     Datacenter = "eu"
)

type Config struct {
	*oauth2.Config
}

func GetEndpoint(dc Datacenter) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:   "https://accounts.zoho." + string(dc) + "/oauth/v2/auth",
		TokenURL:  "https://accounts.zoho." + string(dc) + "/oauth/v2/token",
		AuthStyle: oauth2.AuthStyleInParams,
	}
}

func (c *Config) GetAuthUrl(state string) string {
	return c.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (c *Config) GetToken(code string) *oauth2.Token {
	cxt := context.Background()
	tok, err := c.Exchange(cxt, code)

	if err != nil {
		panic(err)
	}
	return tok
}
