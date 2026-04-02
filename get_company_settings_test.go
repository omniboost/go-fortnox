package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetCompanySettings(t *testing.T) {
	req := client.NewGetCompanySettingsRequest()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
