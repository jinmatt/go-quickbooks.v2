package quickbooks

import (
	"encoding/json"
	"fmt"
)

// InvoiceObject the complete quickbooks invoice object type
type InvoiceObject struct {
	Invoice Invoice `json:"Invoice"`
	Time    string  `json:"time"`
}

// Invoice quickbooks invoice type
type Invoice struct {
	ID           string         `json:"Id,omitempty"`
	Deposit      int            `json:"Deposit,omitempty"`
	Domain       string         `json:"domain,omitempty"`
	Sparse       bool           `json:"sparse,omitempty"`
	SyncToken    string         `json:"SyncToken,omitempty"`
	CustomField  *[]CustomField `json:"CustomField,omitempty"`
	DocNumber    string         `json:"DocNumber,omitempty"`
	TxnDate      string         `json:"TxnDate,omitempty"`
	LinkedTxn    *[]LinkedTxn   `json:"LinkedTxn,omitempty"`
	Line         []InvoiceLine  `json:"Line"`
	TxnTaxDetail struct {
		TxnTaxCodeRef *TaxCodeRef `json:"TxnTaxCodeRef,omitempty"`
		TotalTax      float64     `json:"TotalTax"`
		TaxLine       []TaxLine   `json:"TaxLine,omitempty"`
	} `json:"TxnTaxDetail,omitempty"`
	CustomerRef  *CustomerRef `json:"CustomerRef"`
	CustomerMemo struct {
		Value string `json:"value"`
	} `json:"CustomerMemo,omitempty"`
	BillAddr     *Address `json:"BillAddr"`
	ShipAddr     *Address `json:"ShipAddr"`
	SalesTermRef struct {
		Value string `json:"value"`
	} `json:"SalesTermRef,omitempty"`
	DueDate               string  `json:"DueDate,omitempty"`
	TotalAmt              float64 `json:"TotalAmt,omitempty"`
	ApplyTaxAfterDiscount bool    `json:"ApplyTaxAfterDiscount,omitempty"`
	PrintStatus           string  `json:"PrintStatus,omitempty"`
	EmailStatus           string  `json:"EmailStatus,omitempty"`
	BillEmail             struct {
		Address string `json:"Address"`
	} `json:"BillEmail,omitempty"`
	Balance  float64 `json:"Balance,omitempty"`
	MetaData struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// InvoiceLine quickbooks invoice line item object
type InvoiceLine struct {
	ID                  string               `json:"Id,omitempty"`
	LineNum             int                  `json:"LineNum,omitempty"`
	Description         string               `json:"Description,omitempty"`
	Amount              float64              `json:"Amount"`
	DetailType          string               `json:"DetailType"`
	SalesItemLineDetail *SalesItemLineDetail `json:"SalesItemLineDetail,omitempty"`
	SubTotalLineDetail  interface{}          `json:"SubTotalLineDetail,omitempty"`
}

// CustomField quickbooks invoice customer field object
type CustomField struct {
	DefinitionID string `json:"DefinitionId"`
	Name         string `json:"Name,omitempty"`
	Type         string `json:"Type"`
	StringValue  string `json:"StringValue,omitempty"`
}

type LinkedTxn struct {
	TxnID   string `json:"TxnId"`
	TxnType string `json:"TxnType"`
}

type TaxCodeRef struct {
	Value string `json:"value"`
}

// SalesItemLineDetail quickbooks invoice sales line item details object
type SalesItemLineDetail struct {
	ItemRef    *ItemRef    `json:"ItemRef"`
	UnitPrice  int         `json:"UnitPrice"`
	Qty        int         `json:"Qty"`
	TaxCodeRef *TaxCodeRef `json:"TaxCodeRef,omitempty"`
}

// TaxLine quickbooks invoice tax line item object
type TaxLine struct {
	Amount        float64 `json:"Amount"`
	DetailType    string  `json:"DetailType,omitempty"`
	TaxLineDetail *struct {
		TaxRateRef       *TaxCodeRef `json:"TaxRateRef,omitempty"`
		PercentBased     bool        `json:"PercentBased"`
		TaxPercent       int         `json:"TaxPercent"`
		NetAmountTaxable float64     `json:"NetAmountTaxable"`
	} `json:"TaxLineDetail,omitempty"`
}

// CreateInvoice creates an invoice on quickbooks
func (q *Quickbooks) CreateInvoice(invoice Invoice) (*InvoiceObject, error) {
	endpoint := fmt.Sprintf("/company/%s/invoice", q.RealmID)

	res, err := q.makePostRequest(endpoint, invoice)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newInvoice := InvoiceObject{}
	err = json.NewDecoder(res.Body).Decode(&newInvoice)
	if err != nil {
		return nil, err
	}

	return &newInvoice, nil
}
