package quickbooks

import (
	"encoding/json"
	"fmt"
)

// ItemObject the complete quickbooks item object type
type ItemObject struct {
	Item Item   `json:"Item"`
	Time string `json:"time"`
}

// Item quickbooks item type
type Item struct {
	ID                 string      `json:"Id,omitempty"`
	Name               string      `json:"Name"`
	Description        string      `json:"Description,omitempty"`
	Type               string      `json:"Type"`
	Active             bool        `json:"Active,omitempty"`
	FullyQualifiedName string      `json:"FullyQualifiedName,omitempty"`
	Taxable            bool        `json:"Taxable,omitempty"`
	UnitPrice          int         `json:"UnitPrice,omitempty"`
	IncomeAccountRef   *AccountRef `json:"IncomeAccountRef"`
	PurchaseDesc       string      `json:"PurchaseDesc,omitempty"`
	PurchaseCost       int         `json:"PurchaseCost,omitempty"`
	ExpenseAccountRef  *AccountRef `json:"ExpenseAccountRef,omitempty"`
	AssetAccountRef    *AccountRef `json:"AssetAccountRef,omitempty"`
	TrackQtyOnHand     bool        `json:"TrackQtyOnHand,omitempty"`
	QtyOnHand          int         `json:"QtyOnHand,omitempty"`
	InvStartDate       string      `json:"InvStartDate,omitempty"`
	Domain             string      `json:"domain,omitempty"`
	Sparse             bool        `json:"sparse,omitempty"`
	SyncToken          string      `json:"SyncToken,omitempty"`
	MetaData           *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// ItemRef quickbooks item reference object
type ItemRef struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}

// CreateItem creates an item on quickbooks
func (q *Quickbooks) CreateItem(item Item) (*ItemObject, error) {
	endpoint := fmt.Sprintf("/company/%s/item", q.RealmID)

	res, err := q.makePostRequest(endpoint, item)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newItem := ItemObject{}
	err = json.NewDecoder(res.Body).Decode(&newItem)
	if err != nil {
		return nil, err
	}

	return &newItem, nil
}
