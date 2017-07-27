package quickbooks

import (
	"encoding/json"
	"fmt"
)

// Company quickbooks object type
type Company struct {
	CompanyInfo struct {
		CompanyName string `json:"CompanyName"`
		LegalName   string `json:"LegalName"`
		CompanyAddr struct {
			ID                     string `json:"Id"`
			Line1                  string `json:"Line1"`
			City                   string `json:"City"`
			Country                string `json:"Country"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			PostalCode             string `json:"PostalCode"`
		} `json:"CompanyAddr"`
		CustomerCommunicationAddr struct {
			ID                     string `json:"Id"`
			Line1                  string `json:"Line1"`
			City                   string `json:"City"`
			Country                string `json:"Country"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			PostalCode             string `json:"PostalCode"`
		} `json:"CustomerCommunicationAddr"`
		LegalAddr struct {
			ID                     string `json:"Id"`
			Line1                  string `json:"Line1"`
			City                   string `json:"City"`
			Country                string `json:"Country"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			PostalCode             string `json:"PostalCode"`
		} `json:"LegalAddr"`
		PrimaryPhone struct {
			FreeFormNumber string `json:"FreeFormNumber"`
		} `json:"PrimaryPhone"`
		CompanyStartDate     string `json:"CompanyStartDate"`
		FiscalYearStartMonth string `json:"FiscalYearStartMonth"`
		Country              string `json:"Country"`
		Email                struct {
			Address string `json:"Address"`
		} `json:"Email"`
		WebAddr struct {
		} `json:"WebAddr"`
		SupportedLanguages string `json:"SupportedLanguages"`
		NameValue          []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"NameValue"`
		Domain    string `json:"domain"`
		Sparse    bool   `json:"sparse"`
		ID        string `json:"Id"`
		SyncToken string `json:"SyncToken"`
		MetaData  struct {
			CreateTime      string `json:"CreateTime"`
			LastUpdatedTime string `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"CompanyInfo"`
	Time string `json:"time"`
}

// GetCompanyInfo returns company info bases on realmID/companyID passed to NewClient options
func (q *Quickbooks) GetCompanyInfo() (*Company, error) {
	endpoint := fmt.Sprintf("/company/%s/companyinfo/%s", q.RealmID, q.RealmID)

	res, err := q.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	company := Company{}
	err = json.NewDecoder(res.Body).Decode(&company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}
