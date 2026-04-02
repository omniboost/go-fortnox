package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetCustomers(t *testing.T) {
	req := client.NewGetCustomersRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
