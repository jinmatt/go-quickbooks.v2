package test

import (
	"testing"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/jinmatt/go-quickbooks.v2"
	seed "github.com/jinmatt/go-seed-rand"
	"github.com/tylerb/is"
)

func TestCreateCustomer(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, true)

	customer := quickbooks.Customer{}

	firstName := randomdata.FirstName(randomdata.RandomGender)
	lastName := randomdata.LastName() + seed.RandomKey(7)

	customer.GivenName = firstName
	customer.FamilyName = lastName
	customer.DisplayName = firstName + " " + lastName

	newCustomer, err := qbo.CreateCustomer(customer)
	is.NotErr(err)
	is.NotNil(newCustomer.Customer.ID)
	is.Equal(customer.DisplayName, newCustomer.Customer.DisplayName)
}
