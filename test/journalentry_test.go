package test

import (
	"testing"
	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/jinmatt/go-quickbooks.v2"
	"github.com/tylerb/is"
	"strconv"
)

func TestCreateJE(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(RealmID, AccessToken, true)

	random_index := randomdata.Number(1,5)
	amount_sum := 0.00
	journalEntryLines := []quickbooks.Line{}
	for random_index > 0 {
		amount := float64(randomdata.Number(20))
		journalEntryLines = append(journalEntryLines,createJournalEntryLine(strconv.Itoa(random_index), randomdata.SillyName(), float64(amount), "Debit", "8", "Bank Charges"))
		amount_sum += amount
		random_index = random_index - 1
	}

	if random_index == 0 {
			journalEntryLines = append(journalEntryLines,createJournalEntryLine(strconv.Itoa(random_index), randomdata.SillyName(), float64(amount_sum), "Credit", "35", "Checking"))
	}

	journalEntry := quickbooks.Journalentry{
    	Line: journalEntryLines,
  	}	

	JournalentryObject, err := qbo.CreateJE(journalEntry)
	is.NotNil(JournalentryObject.Journalentry.ID)
	is.NotErr(err)

}

func createJournalEntryLine(lineID string, description string, amount float64, postingType string, accountNum string, accountName string) quickbooks.Line {
	line := quickbooks.Line{
		LineID:      lineID,
		Description: description,
		Amount:      amount,
		DetailType:  "JournalEntryLineDetail",
		JournalEntryLineDetail: &quickbooks.JournalEntryLineDetail{
			PostingType: postingType,
			AccountRef: quickbooks.JournalEntryRef{
				Value: accountNum,
				Name:  accountName,
			},
		},
	}
	return line

}
