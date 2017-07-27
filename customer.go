package quickbooks

import (
	"encoding/json"
	"fmt"
)

// CustomerObject the complete quickbooks customer object type
type CustomerObject struct {
	Customer Customer `json:"Customer"`
	Time     string   `json:"time"`
}

// Customer quickbooks customer type
type Customer struct {
	ID                      string   `json:"Id,omitempty"`
	Taxable                 bool     `json:"Taxable,omitempty"`
	BillAddr                *Address `json:"BillAddr,omitempty"`
	ShipAddr                *Address `json:"ShipAddr,omitempty"`
	Job                     bool     `json:"Job,omitempty"`
	BillWithParent          bool     `json:"BillWithParent,omitempty"`
	Balance                 float64  `json:"Balance,omitempty"`
	BalanceWithJobs         float64  `json:"BalanceWithJobs,omitempty"`
	PreferredDeliveryMethod string   `json:"PreferredDeliveryMethod,omitempty"`
	Domain                  string   `json:"domain,omitempty"`
	Sparse                  bool     `json:"sparse,omitempty"`
	SyncToken               string   `json:"SyncToken,omitempty"`
	GivenName               string   `json:"GivenName"`
	MiddleName              string   `json:"MiddleName,omitempty"`
	FamilyName              string   `json:"FamilyName"`
	FullyQualifiedName      string   `json:"FullyQualifiedName,omitempty"`
	CompanyName             string   `json:"CompanyName,omitempty"`
	DisplayName             string   `json:"DisplayName"`
	PrintOnCheckName        string   `json:"PrintOnCheckName,omitempty"`
	Active                  bool     `json:"Active,omitempty"`

	PrimaryPhone *struct {
		FreeFormNumber string `json:"FreeFormNumber"`
	} `json:"PrimaryPhone,omitempty"`
	PrimaryEmailAddr *struct {
		Address string `json:"Address"`
	} `json:"PrimaryEmailAddr,omitempty"`

	MetaData *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// Address quickbooks address object
type Address struct {
	ID                     string `json:"Id,omitempty"`
	Line1                  string `json:"Line1"`
	Line2                  string `json:"Line2,omitempty"`
	City                   string `json:"City"`
	CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
	PostalCode             string `json:"PostalCode"`
	Lat                    string `json:"Lat,omitempty"`
	Long                   string `json:"Long,omitempty"`
}

// CustomerRef quickbooks customer reference object
type CustomerRef struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}

// CreateCustomer creates a customer on quickbooks
func (q *Quickbooks) CreateCustomer(customer Customer) (*CustomerObject, error) {
	endpoint := fmt.Sprintf("/company/%s/customer", q.RealmID)

	res, err := q.makePostRequest(endpoint, customer)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newCustomer := CustomerObject{}
	err = json.NewDecoder(res.Body).Decode(&newCustomer)
	if err != nil {
		return nil, err
	}

	return &newCustomer, nil
}
