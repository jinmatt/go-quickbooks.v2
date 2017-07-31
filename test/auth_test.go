package test

import (
	"net/url"
	"testing"

	quickbooks "github.com/jinmatt/go-quickbooks.v2"
	"github.com/tylerb/is"
)

func TestGetConnectURL(t *testing.T) {
	is := is.New(t)

	connectURL, err := quickbooks.GetConnectURL(ClientID, AccountScope, RedirectURI, "test-token", true)
	is.NotErr(err)

	URL, err := url.Parse(connectURL)
	is.NotErr(err)
	q := URL.Query()
	is.Equal(q.Get("state"), "test-token")
}

func TestGetBearerToken(t *testing.T) {
	t.Skip("Skipping test: authorization code expires after one time call")

	is := is.New(t)

	bearerToken, err := quickbooks.GetBearerToken(ClientID, ClientSecret, AuthCode, RedirectURI, true)
	is.NotErr(err)
	is.NotNil(bearerToken.AccessToken)
	is.NotNil(bearerToken.RefreshToken)
	is.Equal(bearerToken.ExpiresIn, 3600)
}

func TestRefreshToken(t *testing.T) {
	is := is.New(t)

	bearerToken, err := quickbooks.RefreshToken(ClientID, ClientSecret, RefreshToken, true)
	is.NotErr(err)
	is.NotNil(bearerToken.AccessToken)
	is.NotNil(bearerToken.RefreshToken)
	is.Equal(bearerToken.ExpiresIn, 3600)
}
