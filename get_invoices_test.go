package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	fortnox "github.com/omniboost/go-fortnox"
)

func TestGetInvoices(t *testing.T) {
	req := client.NewGetInvoicesRequest()
	req.QueryParams().FromDate = fortnox.Date{time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)}
	req.QueryParams().ToDate = fortnox.Date{time.Date(2023, 2, 4, 0, 0, 0, 0, time.Local)}
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
