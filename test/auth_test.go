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
