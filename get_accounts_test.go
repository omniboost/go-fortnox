package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	req := client.NewGetAccountsRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestAllAccounts(t *testing.T) {
	req := client.NewGetAccountsRequest()
	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
