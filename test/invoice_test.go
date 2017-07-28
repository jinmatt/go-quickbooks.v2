package test

import (
	"testing"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/jinmatt/go-quickbooks.v2"
	"github.com/jinmatt/go-quickbooks.v2/sdk/consts"
	seed "github.com/jinmatt/go-seed-rand"
	"github.com/tylerb/is"
)

func TestCreateInvoice(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, true)

	// create a chart of account for item
	account := quickbooks.Account{}
	account.Name = randomdata.SillyName() + seed.RandomKey(7)
	account.AccountType = consts.QBAccountIncomeType
	newAccount, err := qbo.CreateAccount(account)
	is.NotErr(err)

	// create an item to charge the customer
	item := quickbooks.Item{}
	item.Name = randomdata.SillyName() + seed.RandomKey(7)
	item.IncomeAccountRef = &quickbooks.AccountRef{
		Value: newAccount.Account.ID,
		Name:  newAccount.Account.Name,
	}
	item.Type = consts.QBItemServiceType
	newItem, err := qbo.CreateItem(item)
	is.NotErr(err)

	// create customer
	customer := quickbooks.Customer{}
	firstName := randomdata.FirstName(randomdata.RandomGender)
	lastName := randomdata.LastName() + seed.RandomKey(7)
	customer.GivenName = firstName
	customer.FamilyName = lastName
	customer.DisplayName = firstName + " " + lastName
	newCustomer, err := qbo.CreateCustomer(customer)
	is.NotErr(err)

	// create Invoice to test
	invoice := quickbooks.Invoice{}
	invoice.CustomerRef = &quickbooks.CustomerRef{
		Value: newCustomer.Customer.ID,
	}

	invoiceLine := quickbooks.InvoiceLine{
		Amount:     500,
		DetailType: consts.QBSalesItemLineDetail,
		SalesItemLineDetail: &quickbooks.SalesItemLineDetail{
			ItemRef: &quickbooks.ItemRef{
				Value: newItem.Item.ID,
			},
		},
	}
	invoice.Line = append(invoice.Line, invoiceLine)

	newInvoice, err := qbo.CreateInvoice(invoice)
	is.NotErr(err)
	is.NotNil(newInvoice.Invoice.ID)
}
