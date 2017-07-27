package test

import (
	"testing"

	"github.com/jinmatt/go-quickbooks.v2"
	"github.com/tylerb/is"
)

func TestCompanyInfo(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, RefreshToken, true)

	company, err := qbo.GetCompanyInfo()
	is.NotErr(err)
	is.NotNil(company.CompanyInfo.CompanyName)
}
