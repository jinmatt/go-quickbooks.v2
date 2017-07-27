package quickbooks

// InvoiceObject the complete quickbooks invoice object type
type InvoiceObject struct {
	Invoice Invoice `json:"Invoice"`
	Time    string  `json:"time"`
}

// Invoice quickbooks invoice type
type Invoice struct {
	Deposit   int    `json:"Deposit"`
	Domain    string `json:"domain"`
	Sparse    bool   `json:"sparse"`
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
	MetaData  struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData"`
	CustomField []struct {
		DefinitionID string `json:"DefinitionId"`
		Name         string `json:"Name"`
		Type         string `json:"Type"`
		StringValue  string `json:"StringValue"`
	} `json:"CustomField"`
	DocNumber string `json:"DocNumber"`
	TxnDate   string `json:"TxnDate"`
	LinkedTxn []struct {
		TxnID   string `json:"TxnId"`
		TxnType string `json:"TxnType"`
	} `json:"LinkedTxn"`
	Line []struct {
		ID                  string  `json:"Id,omitempty"`
		LineNum             int     `json:"LineNum,omitempty"`
		Description         string  `json:"Description,omitempty"`
		Amount              float64 `json:"Amount"`
		DetailType          string  `json:"DetailType"`
		SalesItemLineDetail struct {
			ItemRef struct {
				Value string `json:"value"`
				Name  string `json:"name"`
			} `json:"ItemRef"`
			UnitPrice  int `json:"UnitPrice"`
			Qty        int `json:"Qty"`
			TaxCodeRef struct {
				Value string `json:"value"`
			} `json:"TaxCodeRef"`
		} `json:"SalesItemLineDetail,omitempty"`
		SubTotalLineDetail struct {
		} `json:"SubTotalLineDetail,omitempty"`
	} `json:"Line"`
	TxnTaxDetail struct {
		TxnTaxCodeRef struct {
			Value string `json:"value"`
		} `json:"TxnTaxCodeRef"`
		TotalTax float64 `json:"TotalTax"`
		TaxLine  []struct {
			Amount        float64 `json:"Amount"`
			DetailType    string  `json:"DetailType"`
			TaxLineDetail struct {
				TaxRateRef struct {
					Value string `json:"value"`
				} `json:"TaxRateRef"`
				PercentBased     bool    `json:"PercentBased"`
				TaxPercent       int     `json:"TaxPercent"`
				NetAmountTaxable float64 `json:"NetAmountTaxable"`
			} `json:"TaxLineDetail"`
		} `json:"TaxLine"`
	} `json:"TxnTaxDetail"`
	CustomerRef struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	} `json:"CustomerRef"`
	CustomerMemo struct {
		Value string `json:"value"`
	} `json:"CustomerMemo"`
	BillAddr struct {
		ID    string `json:"Id"`
		Line1 string `json:"Line1"`
		Line2 string `json:"Line2"`
		Line3 string `json:"Line3"`
		Line4 string `json:"Line4"`
		Lat   string `json:"Lat"`
		Long  string `json:"Long"`
	} `json:"BillAddr"`
	ShipAddr struct {
		ID                     string `json:"Id"`
		Line1                  string `json:"Line1"`
		City                   string `json:"City"`
		CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
		PostalCode             string `json:"PostalCode"`
		Lat                    string `json:"Lat"`
		Long                   string `json:"Long"`
	} `json:"ShipAddr"`
	SalesTermRef struct {
		Value string `json:"value"`
	} `json:"SalesTermRef"`
	DueDate               string  `json:"DueDate"`
	TotalAmt              float64 `json:"TotalAmt"`
	ApplyTaxAfterDiscount bool    `json:"ApplyTaxAfterDiscount"`
	PrintStatus           string  `json:"PrintStatus"`
	EmailStatus           string  `json:"EmailStatus"`
	BillEmail             struct {
		Address string `json:"Address"`
	} `json:"BillEmail"`
	Balance float64 `json:"Balance"`
}
