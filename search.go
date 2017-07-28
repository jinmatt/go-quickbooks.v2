package quickbooks

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// QueryResponse quickbooks search response type
type QueryResponse map[string]interface{}

// Search quickbooks document with a query string
func (q *Quickbooks) Search(query string) (*QueryResponse, error) {
	endpoint := fmt.Sprintf("/company/%s/query?query=%s", q.RealmID, url.QueryEscape(query))

	res, err := q.makePostRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	queryResponse := QueryResponse{}
	err = json.NewDecoder(res.Body).Decode(&queryResponse)
	if err != nil {
		return nil, err
	}

	return &queryResponse, nil
}

// ToAccount converts a search QueryRespose to Account array type
func (qr QueryResponse) ToAccount() ([]Account, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Account"])
	if err != nil {
		return nil, err
	}

	accounts := []Account{}
	err = json.Unmarshal(b, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// ToItem converts a search QueryRespose to Item array type
func (qr QueryResponse) ToItem() ([]Item, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Item"])
	if err != nil {
		return nil, err
	}

	items := []Item{}
	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// ToCustomer converts a search QueryRespose to Customer array type
func (qr QueryResponse) ToCustomer() ([]Customer, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Customer"])
	if err != nil {
		return nil, err
	}

	customers := []Customer{}
	err = json.Unmarshal(b, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// ToInvoice converts a search QueryRespose to Invoice array type
func (qr QueryResponse) ToInvoice() ([]Invoice, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Invoice"])
	if err != nil {
		return nil, err
	}

	invoices := []Invoice{}
	err = json.Unmarshal(b, &invoices)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// ToPayment converts a search QueryRespose to Payment array type
func (qr QueryResponse) ToPayment() ([]Payment, error) {
	queryResponse := qr["QueryResponse"]
	document := queryResponse.(map[string]interface{})

	b, err := json.Marshal(document["Payment"])
	if err != nil {
		return nil, err
	}

	payments := []Payment{}
	err = json.Unmarshal(b, &payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
