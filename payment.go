package quickbooks

import (
	"encoding/json"
	"fmt"
)

// PaymentObject the complete quickbooks payment object type
type PaymentObject struct {
	Payment Payment `json:"Payment"`
	Time    string  `json:"time"`
}

// Payment quickbooks payment type
type Payment struct {
	ID                  string        `json:"Id,omitempty"`
	CustomerRef         *CustomerRef  `json:"CustomerRef,omitempty"`
	DepositToAccountRef *AccountRef   `json:"DepositToAccountRef,omitempty"`
	TotalAmt            float64       `json:"TotalAmt"`
	UnappliedAmt        float64       `json:"UnappliedAmt,omitempty"`
	ProcessPayment      bool          `json:"ProcessPayment,omitempty"`
	Domain              string        `json:"domain,omitempty"`
	Sparse              bool          `json:"sparse,omitempty"`
	SyncToken           string        `json:"SyncToken,omitempty"`
	TxnDate             string        `json:"TxnDate,omitempty"`
	Line                []PaymentLine `json:"Line"`
	MetaData            *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// PaymentLine quickbooks payment line object
type PaymentLine struct {
	Amount    float64     `json:"Amount"`
	LinkedTxn []LinkedTxn `json:"LinkedTxn"`
	LineEx    *LineEx     `json:"LineEx,omitempty"`
}

// LineEx quickbooks payment LineEx object
type LineEx struct {
	Any []struct {
		Name         string `json:"name"`
		DeclaredType string `json:"declaredType"`
		Scope        string `json:"scope"`
		Value        struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"value"`
		Nil             bool `json:"nil"`
		GlobalScope     bool `json:"globalScope"`
		TypeSubstituted bool `json:"typeSubstituted"`
	} `json:"any"`
}

// CreatePayment creates a payment on quickbooks
func (q *Quickbooks) CreatePayment(payment Payment) (*PaymentObject, error) {
	endpoint := fmt.Sprintf("/company/%s/payment", q.RealmID)

	res, err := q.makePostRequest(endpoint, payment)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newPayment := PaymentObject{}
	err = json.NewDecoder(res.Body).Decode(&newPayment)
	if err != nil {
		return nil, err
	}

	return &newPayment, nil
}
