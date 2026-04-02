package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetOrders(t *testing.T) {
	req := client.NewGetOrdersRequest()
	// req.QueryParams().OrderNumber = "TEST-002"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
