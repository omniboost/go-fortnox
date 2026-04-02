package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestPostOrder(t *testing.T) {
	req := client.NewPostOrderRequest()
	order := req.RequestBody().Order
	order.OrderType = "Test from API #2"
	req.RequestBody().Order = order
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
