package quickbooks

import (
	"encoding/json"
	"fmt"
)

// JournalentryObject the complete quickbooks journal entry object type
type JournalentryObject struct {
	Journalentry Journalentry `json:"Journalentry"`
	Time         string       `json:"time"`
}

// Journal Entry quickbooks Journal Entry type
type Journalentry struct {
	ID         string    `json:"Id,omitempty"`
	Adjustment bool      `json:"Adjustment,omitempty"`
	Domain     string    `json:"domain,omitempty"`
	Sparse     bool      `json:"sparse,omitempty"`
	SyncToken  string    `json:"SyncToken,omitempty"`
	TxnDate    string    `json:"TxnDate,omitempty"`
	Line       []Line    `json:"Line"`
	MetaData   *MetaData `json:"MetaData,omitempty"`
}

// Line type - part of Journal Entry
type Line struct {
	LineID                 string                  `json:"Id,omitempty"`
	Description            string                  `json:"Description,omitempty"`
	Amount                 float64                 `json:"Amount,omitempty"`
	DetailType             string                  `json:"DetailType,omitempty"`
	JournalEntryLineDetail *JournalEntryLineDetail `json:"JournalEntryLineDetail,omitempty"`
}

// JournalEntryLineDetail - part of Journal Entry
type JournalEntryLineDetail struct {
	AccountRef  JournalEntryRef `json:"AccountRef,omitempty"`
	PostingType string          `json:"PostingType,omitempty"`
}

// Metadata - info about when the journal entry was created/updated.
type MetaData struct {
	CreateTime      string `json:"CreateTime, omitempty"`
	LastUpdatedTime string `json:"LastUpdatedTime, omitempty"`
}

type JournalEntryRef struct {
	Value string `json:"value, omitempty"`
	Name  string `json:"name, omitempty"`
}

// CreateJE creates a journal entry on quickbooks
func (q *Quickbooks) CreateJE(journalentry Journalentry) (*JournalentryObject, error) {
	endpoint := fmt.Sprintf("/company/%s/journalentry", q.RealmID)

	response, err := q.makePostRequest(endpoint, journalentry)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	newJE := JournalentryObject{}
	err = json.NewDecoder(response.Body).Decode(&newJE)
	if err != nil {
		return nil, err
	}

	return &newJE, nil
}
