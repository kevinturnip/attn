package model

import (
	"testing"

	"log"
)

func TestNewCurrencies(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var body Currencies
	body.Id = 999
	body.Name = "name999"
	curr, err := MockCreateCurrency(body.Name, body.Id)
	if curr.Id == 0 {
		t.Error("Failed create new currency", err)
	}
}
