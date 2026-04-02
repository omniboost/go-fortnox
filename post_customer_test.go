package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	fortnox "github.com/omniboost/go-fortnox"
)

func TestPostCustomer(t *testing.T) {
	req := client.NewPostCustomerRequest()
	customer := req.RequestBody().Customer
	customer.Name = "TEST"
	customer.Active = true
	customer.Type = fortnox.CustomerTypePrivate
	customer.VATType = fortnox.VATTypeSEVAT
	req.RequestBody().Customer = customer
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
