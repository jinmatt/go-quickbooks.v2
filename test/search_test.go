package test

import (
	"testing"

	quickbooks "github.com/jinmatt/go-quickbooks.v2"
	"github.com/tylerb/is"
)

func TestSearch(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, true)

	// Query account test
	query := "Select * from Account"
	queryResponse, err := qbo.Search(query)
	is.NotErr(err)

	accounts, err := queryResponse.ToAccount()
	is.NotErr(err)
	is.NotZero(len(accounts))
	is.NotNil(accounts[0].ID)

	// Query item test
	query = "Select * from Item"
	queryResponse, err = qbo.Search(query)
	is.NotErr(err)

	items, err := queryResponse.ToItem()
	is.NotErr(err)
	is.NotZero(len(items))
	is.NotNil(items[0].ID)

	// Query customer test
	query = "Select * from Customer"
	queryResponse, err = qbo.Search(query)
	is.NotErr(err)

	customers, err := queryResponse.ToCustomer()
	is.NotErr(err)
	is.NotZero(len(customers))
	is.NotNil(customers[0].ID)

	// Query invoice test
	query = "Select * from Invoice"
	queryResponse, err = qbo.Search(query)
	is.NotErr(err)

	invoices, err := queryResponse.ToInvoice()
	is.NotErr(err)
	is.NotZero(len(invoices))
	is.NotNil(invoices[0].ID)

	// Query payment test
	query = "Select * from Payment"
	queryResponse, err = qbo.Search(query)
	is.NotErr(err)

	payments, err := queryResponse.ToPayment()
	is.NotErr(err)
	is.NotZero(len(payments))
	is.NotNil(payments[0].ID)
}
