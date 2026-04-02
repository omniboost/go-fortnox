package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetPredefinedVoucherSeries(t *testing.T) {
	req := client.NewGetPredefinedVoucherSeriesRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
