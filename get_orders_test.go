package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetOrders(t *testing.T) {
	req := client.NewGetOrdersRequest()
	// req.QueryParams().OrderNumber = "TEST-002"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
