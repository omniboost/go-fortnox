package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetInvoices(t *testing.T) {
	req := client.NewGetInvoicesRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
